package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "go/token"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
	_ "github.com/lib/pq"
	// "golang.org/x/crypto/bcrypt"
	// "github.com/dgrijalva/jwt-go"
)

type Employee struct{
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	MiddleName string `json:"middleName"`
	EmailID string `json:"emailid"`
	SapID string `json:"sapid"`
	Role string `json:"role"`
	Organization string `json:"organization"`
	ReportingTo string `json:"reportingto"`
	PhoneNumber string `json:"phonenumber"`
}

type Training struct{
	Category string `json:"category"`
	CourseName string `json:"courseName"`
	CourseCode string `json:"courseCode"`
}

//Check Error Function
func checkErr(err error){
	if err != nil{
		panic(err)
	}
}

var db *sql.DB
var err error

func setupDB() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	db, err = sql.Open("postgres", dbinfo)

	checkErr(err)

	return db
}


func getEnvVars(){
	err := godotenv.Load("credentials.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

//JWT Authentication
// var SECRET_KEY = []byte("gosecretkey")

// func getHash(pwd []byte) string{
// 	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
// 	checkErr(err)
// 	return string(hash)
// }

// func GenerateJWT()(string, error){
// 	token := jwt.New(jwt.SigningMethodHS256)
// 	tokenString, err := token.SignedString(SECRET_KEY)
// 	checkErr(err)
// 	return tokenString, nil
// }


//Main Function
func main(){
	getEnvVars()
	setupDB()
	defer db.Close()
	router := mux.NewRouter()

	//Employee Routes

	router.HandleFunc("/employees", GetEmployees).Methods("GET")

	router.HandleFunc("/employees/{emailid}", GetEmployee).Methods("GET")

	router.HandleFunc("/employees", CreateEmployee).Methods("POST")

	router.HandleFunc("/employees/{emailid}", UpdateEmployee).Methods("PUT")

	router.HandleFunc("/employees/{emailid}", DeleteEmployee).Methods("DELETE")

	//Training Routes
	
	router.HandleFunc("/trainings", GetTrainings).Methods("GET")

	router.HandleFunc("/trainings/{coursecode}", GetTraining).Methods("GET")

	router.HandleFunc("/trainings/{coursecode}", DeleteTraining).Methods("DELETE")

	router.HandleFunc("/trainings/{coursecode}", UpdateTraining).Methods("PUT")

	router.HandleFunc("/trainings", CreateTraining).Methods("POST")

	fmt.Println("Listening at port 9080")
	log.Fatal(http.ListenAndServe(":9080", &CORSRouterDecorator{router}))
}

//EMPLOYEE METHODS//

//Get Single Employee
func GetEmployee(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result, err := db.Query("SELECT * FROM employees WHERE email_id = $1", params["emailid"])
	checkErr(err)
	defer result.Close()
	var employee Employee
	for result.Next(){
		err := result.Scan(&employee.FirstName, &employee.LastName, &employee.MiddleName, &employee.EmailID, &employee.SapID, &employee.Role, &employee.Organization, &employee.ReportingTo, &employee.PhoneNumber)
		checkErr(err)
	}
	json.NewEncoder(w).Encode(employee)
}

//Get All Employees
func GetEmployees(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var employees []Employee
	result, err := db.Query("SELECT * FROM employees")
	checkErr(err)
	defer result.Close()
	for result.Next() {
		var employee Employee
		err := result.Scan(&employee.FirstName, &employee.LastName, &employee.MiddleName, &employee.EmailID, &employee.SapID, &employee.Role, &employee.Organization, &employee.ReportingTo, &employee.PhoneNumber)
		checkErr(err)
		employees = append(employees, employee)
	}
	json.NewEncoder(w).Encode(employees)
}

//Create Employee
func CreateEmployee(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	stmt, err := db.Prepare("INSERT INTO employees(first_name, last_name, middle_name, email_id, sap_id, role, organization, reporting_to, phone_number) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)")
	checkErr(err)
	body, err := ioutil.ReadAll(r.Body)
	checkErr(err)
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	first_name := keyVal["firstName"]
	last_name := keyVal["lastName"]
	middle_name := keyVal["middleName"]
	email_id := keyVal["emailid"]
	sap_id := keyVal["sapid"]
	role := keyVal["role"]
	organization := keyVal["organization"]
	reporting_to := keyVal["reportingto"]
	phone_number := keyVal["phonenumber"]
	_, err = stmt.Exec(first_name, last_name, middle_name, email_id, sap_id, role, organization, reporting_to, phone_number)
	checkErr(err)
	fmt.Fprintln(w, "New Employee Was Created")
}

//Delete Employee
func DeleteEmployee(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    stmt, err := db.Prepare("DELETE FROM employees WHERE email_id = $1")
    checkErr(err)
    _, err = stmt.Exec(params["emailid"])
    checkErr(err)
    fmt.Fprintf(w, "Employee with ID = %s was deleted", params["emailid"])
}

//Update Employee
func UpdateEmployee(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	stmt, err := db.Prepare("UPDATE employees SET first_name = $1, last_name = $2, middle_name = $3, email_id = $4, sap_id = $5, role = $6, organization = $7, reporting_to = $8, phone_number = $9 WHERE email_id = $10")
	checkErr(err)
	body, err := ioutil.ReadAll(r.Body)
	checkErr(err)
	keyVal := make(map[string]string)
    json.Unmarshal(body, &keyVal)
    first_name := keyVal["firstName"]
    last_name := keyVal["lastName"]
    middle_name := keyVal["middleName"]
	email_id := keyVal["emailid"]
	sap_id := keyVal["sapid"]
	role := keyVal["role"]
	organization := keyVal["organization"]
	reporting_to := keyVal["reportingto"]
	phone_number := keyVal["phonenumber"]
    _, err = stmt.Exec(first_name, last_name, middle_name, email_id, sap_id, role, organization, reporting_to, phone_number, params["emailid"])
    checkErr(err)
    fmt.Fprintf(w, "User with ID = %s was updated", params["emailid"])
}

//TRAINING METHODS//

//Get All Trainings
func GetTrainings(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var trainings []Training
	result, err := db.Query("SELECT * FROM trainings")
	checkErr(err)
	defer result.Close()
	for result.Next() {
		var training Training
		err = result.Scan(&training.Category, &training.CourseName, &training.CourseCode)
		checkErr(err)
		trainings = append(trainings, training)
	}
	json.NewEncoder(w).Encode(trainings)
}

//Get Single Training
func GetTraining(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result, err := db.Query("SELECT * FROM trainings WHERE course_code = $1", params["coursecode"])
	checkErr(err)
	defer result.Close()
	var training Training
	for result.Next(){
		err := result.Scan(&training.Category, &training.CourseName, &training.CourseCode)
		checkErr(err)
	}
	json.NewEncoder(w).Encode(training)
}

//Create Training
func CreateTraining(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	stmt, err := db.Prepare("INSERT INTO trainings(category, course_name, course_code) VALUES ($1, $2 , $3)")
	checkErr(err)
	body, err := ioutil.ReadAll(r.Body)
	checkErr(err)
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	category := keyVal["category"]
	course_name := keyVal["courseName"]
	course_code := keyVal["courseCode"]
	_, err = stmt.Exec(category, course_name, course_code)
	checkErr(err)
	fmt.Fprintln(w, "New Training Has Been Created")
}	

//Update Training
func UpdateTraining(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	stmt, err := db.Prepare("UPDATE trainings SET category = $1, course_name = $2 WHERE course_code = $3")
	checkErr(err)
	body, err := ioutil.ReadAll(r.Body)
	checkErr(err)
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	category := keyVal["category"]
	course_name := keyVal["courseName"]
	_, err = stmt.Exec(category, course_name, params["coursecode"])
	checkErr(err)
	fmt.Fprintf(w, "Training with Course Code = %s was Updated", params["coursecode"])
}

//Delete Training
func DeleteTraining(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	stmt, err := db.Prepare("DELETE FROM trainings WHERE course_code = $1")
	checkErr(err)
	_, err = stmt.Exec(params["coursecode"])
	checkErr(err)
	fmt.Fprintf(w, "Training with Course Code = %s Has Been Deleted", params["coursecode"])
}

// CORSRouterDecorator applies CORS headers to a mux.Router
type CORSRouterDecorator struct {
    R *mux.Router
}

func (c *CORSRouterDecorator) ServeHTTP(rw http.ResponseWriter,
    req *http.Request) {
    if origin := req.Header.Get("Origin"); origin != "" {
        rw.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
        rw.Header().Set("Access-Control-Allow-Methods",
            "POST, GET, OPTIONS, PUT, DELETE")
        rw.Header().Set("Access-Control-Allow-Headers",
            "Accept, Accept-Language,"+
                " Content-Type, YourOwnHeader")
    }
    // Stop here if its Preflighted OPTIONS request
    if req.Method == "OPTIONS" {
        return
    }

    c.R.ServeHTTP(rw, req)
}