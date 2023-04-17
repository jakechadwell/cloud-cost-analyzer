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
	EmployeeID string `json:"employeeid"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Dept string `json:"dept"`
	Cloud string `json:"cloud"`
	TrainingAttended string `json:"trainingAttended"`
	TrainingPath string `json:"trainingPath"`
	Email string `json:"email"`
	Infographics string `json:"infographics"`
}

type Training struct{
	TrainingID string `json:"trainingid"`
	CloudID string `json:"cloudid"`
	TrainingDetail string `json:"trainingDetail"`
	TrainingPath string `json:"trainingPath"`
	TrainingDates string `json:"trainingDates"`
}

type Attribute struct{
	AttributeID string `json:"attributeid"`
	AttributeName string `json:"attributeName"`
	Stage string `json:"stage"`
	AttributeType string `json:"attributeType"`
	Count string `json:"count"`
	Course string `json:"course"`
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

	router.HandleFunc("/employees/{employeeid}", GetEmployee).Methods("GET")

	router.HandleFunc("/employees", CreateEmployee).Methods("POST")

	router.HandleFunc("/employees/{employeeid}", UpdateEmployee).Methods("PUT")

	router.HandleFunc("/employees/{employeeid}", DeleteEmployee).Methods("DELETE")

	//Training Routes
	
	router.HandleFunc("/trainings", GetTrainings).Methods("GET")

	router.HandleFunc("/trainings/{trainingid}", GetTraining).Methods("GET")

	router.HandleFunc("/trainings/{trainingid}", DeleteTraining).Methods("DELETE")

	router.HandleFunc("/trainings/{trainingid}", UpdateTraining).Methods("PUT")

	router.HandleFunc("/trainings", CreateTraining).Methods("POST")

	
	fmt.Println("Listening at port 9080")
	log.Fatal(http.ListenAndServe(":9080", &CORSRouterDecorator{router}))
}

//EMPLOYEE METHODS//

//Get Single Employee
func GetEmployee(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result, err := db.Query("SELECT * FROM employees WHERE employee_id = $1", params["employeeid"])
	checkErr(err)
	defer result.Close()
	var employee Employee
	for result.Next(){
		err := result.Scan(&employee.EmployeeID, &employee.FirstName, &employee.LastName, &employee.Dept, &employee.Cloud, &employee.TrainingAttended, &employee.TrainingPath, &employee.Email, &employee.Infographics)
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
		err := result.Scan(&employee.EmployeeID, &employee.FirstName, &employee.LastName, &employee.Dept, &employee.Cloud, &employee.TrainingAttended, &employee.TrainingPath, &employee.Email, &employee.Infographics)
		checkErr(err)
		employees = append(employees, employee)
	}
	json.NewEncoder(w).Encode(employees)
}

//Create Employee
func CreateEmployee(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	stmt, err := db.Prepare("INSERT INTO employees(employee_id, first_name, last_name, dept, cloud, training_attended, training_path, email, infographics) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)")
	checkErr(err)
	body, err := ioutil.ReadAll(r.Body)
	checkErr(err)
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	employee_id := keyVal["employeeid"]
	first_name := keyVal["firstName"]
	last_name := keyVal["lastName"]
	dept := keyVal["dept"]
	cloud := keyVal["cloud"]
	training_attended := keyVal["trainingAttended"]
	training_path := keyVal["trainingPath"]
	email := keyVal["email"]
	infographics := keyVal["infographics"]
	_, err = stmt.Exec(employee_id, first_name, last_name, dept, cloud, training_attended, training_path, email, infographics)
	checkErr(err)
	fmt.Fprintln(w, "New Employee Was Created")
}

//Delete Employee
func DeleteEmployee(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    stmt, err := db.Prepare("DELETE FROM employees WHERE employee_id = $1")
    checkErr(err)
    _, err = stmt.Exec(params["employeeid"])
    checkErr(err)
    fmt.Fprintf(w, "Employee with ID = %s was deleted", params["employeeid"])
}

//Update Employee
func UpdateEmployee(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	stmt, err := db.Prepare("UPDATE employees SET first_name = $1, last_name = $2, dept = $3, cloud = $4, training_attended = $5, training_path = $6, email = $7, infographics = $8 WHERE employee_id = $9")
	checkErr(err)
	body, err := ioutil.ReadAll(r.Body)
	checkErr(err)
	keyVal := make(map[string]string)
    json.Unmarshal(body, &keyVal)
	first_name := keyVal["firstName"]
	last_name := keyVal["lastName"]
	dept := keyVal["dept"]
	cloud := keyVal["cloud"]
	training_attended := keyVal["trainingAttended"]
	training_path := keyVal["trainingPath"]
	email := keyVal["email"]
	infographics := keyVal["infographics"]
    _, err = stmt.Exec(first_name, last_name, dept, cloud, training_attended, training_path, email, infographics, params["employeeid"])
    checkErr(err)
    fmt.Fprintf(w, "User with ID = %s was updated", params["employeeid"])
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
		err = result.Scan(&training.TrainingID, &training.CloudID, &training.TrainingDetail, &training.TrainingPath, &training.TrainingDates)
		checkErr(err)
		trainings = append(trainings, training)
	}
	json.NewEncoder(w).Encode(trainings)
}

//Get Single Training
func GetTraining(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result, err := db.Query("SELECT * FROM trainings WHERE training_id = $1", params["trainingid"])
	checkErr(err)
	defer result.Close()
	var training Training
	for result.Next(){
		err := result.Scan(&training.TrainingID, &training.CloudID, &training.TrainingDetail, &training.TrainingPath, &training.TrainingDates)
		checkErr(err)
	}
	json.NewEncoder(w).Encode(training)
}

//Create Training
func CreateTraining(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	stmt, err := db.Prepare("INSERT INTO trainings(training_id, cloud_id, training_detail, training_path, training_dates) VALUES ($1, $2 , $3, $4, $5)")
	checkErr(err)
	body, err := ioutil.ReadAll(r.Body)
	checkErr(err)
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	training_id := keyVal["trainingid"]
	cloud_id := keyVal["cloudid"]
	training_detail := keyVal["trainingDetail"]
	training_path := keyVal["trainingPath"]
	training_dates := keyVal["trainingDates"]
	_, err = stmt.Exec(training_id, cloud_id, training_detail, training_path, training_dates)
	checkErr(err)
	fmt.Fprintln(w, "New Training Has Been Created")
}	

//Update Training
func UpdateTraining(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	stmt, err := db.Prepare("UPDATE trainings SET cloud_id = $1, training_detail = $2, training_path = $3, training_dates = $4 WHERE training_id = $5")
	checkErr(err)
	body, err := ioutil.ReadAll(r.Body)
	checkErr(err)
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	cloud_id := keyVal["cloudid"]
	training_detail := keyVal["trainingDetail"]
	training_path := keyVal["trainingPath"]
	training_dates := keyVal["trainingDates"]
	_, err = stmt.Exec(cloud_id, training_detail, training_path, training_dates, params["trainingid"])
	checkErr(err)
	fmt.Fprintf(w, "Training with ID = %s was Updated", params["trainingid"])
}

//Delete Training
func DeleteTraining(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	stmt, err := db.Prepare("DELETE FROM trainings WHERE training_id = $1")
	checkErr(err)
	_, err = stmt.Exec(params["trainingid"])
	checkErr(err)
	fmt.Fprintf(w, "Training with ID = %s Has Been Deleted", params["trainingid"])
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