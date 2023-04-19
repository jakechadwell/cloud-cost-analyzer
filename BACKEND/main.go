package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	_ "go/token"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
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

type Cloud struct {
	CloudID                string `json:"cloudid"`                //Could reasonably be an integer instead of string
	CloudTrainingPath      string `json:"cloudtrainingpath"`      //not sure what cloud training path is, set as string for now
	CloudTrainingAvailable string `json:"cloudtrainingavailable"` //Set as string so we can list out multiple trainings
	CloudPointOfContact    string `json:"cloudpointofcontact"`    //Guessing this is actually the person you reach out to about training
	ExternalTrainingPoc    string `json:"externalcontactpoc"`      //Likely the external contact for the training provided, could be saved as a name and/or phone number
}

type Credentials struct {
	Email string `json:"email"`
	Password string `json:"password"`
	Role string `json:"role"`
}

type Token struct {
	Role string `json:"role"`
	Email string `json:"email"`
	TokenString string `json:"tokenString"`
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

// JWT Authentication

func getHash(pwd []byte) string{
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	checkErr(err)
	return string(hash)
}

func GenerateJWT(email, role string)(string, error){
	var SECRET_KEY = []byte("gosecretkey")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(SECRET_KEY)
	checkErr(err)
	return tokenString, nil
}


//Main Function
func main(){
	getEnvVars()
	setupDB()
	defer db.Close()
	router := mux.NewRouter()

	//Employee Routes

	router.HandleFunc("/employees/{email}", GetEmployeeByEmail).Methods("GET")

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

	//Cloud Routes

	router.HandleFunc("/cloud", GetCloudIDs).Methods("GET")

	router.HandleFunc("/cloud/{cloudid}", GetCloudID).Methods("GET")

	router.HandleFunc("/cloud", CreateCloudID).Methods("POST")

	router.HandleFunc("/cloud/{cloudid}", UpdateCloudID).Methods("PUT")

	router.HandleFunc("/cloud/{cloudid}", DeleteCloudID).Methods("DELETE")

	//Credential Routes

	router.HandleFunc("/credentials", GetCredentials).Methods("GET")

	router.HandleFunc("/credentials/{email}", GetCredential).Methods("GET")

	router.HandleFunc("/credentials", CreateCredentials).Methods("POST")

	router.HandleFunc("/credentials/{email}", UpdateCredentials).Methods("PUT")

	router.HandleFunc("/credentials/{email}", DeleteCredentials).Methods("DELETE")

	//Auth Routes

	router.HandleFunc("/signin", SignIn).Methods("POST")

	router.HandleFunc("/signup", SignUp).Methods("POST")

	fmt.Println("Listening at port 9080")
	log.Fatal(http.ListenAndServe(":9080", &CORSRouterDecorator{router}))
}

//EMPLOYEE METHODS//

//Get Single Employee By Email
func GetEmployeeByEmail(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result, err := db.Query("SELECT * FROM employees WHERE email = $1", params["email"])
	checkErr(err)
	defer result.Close()
	var employee Employee
	for result.Next(){
		err := result.Scan(&employee.EmployeeID, &employee.FirstName, &employee.LastName, &employee.Dept, &employee.Cloud, &employee.TrainingAttended, &employee.TrainingPath, &employee.Email, &employee.Infographics)
		checkErr(err)
	}
	json.NewEncoder(w).Encode(employee)
}

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
	IsAuthorized(AdminIndex)
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

//CLOUD METHODS//

// Get All CloudIDs
func GetCloudIDs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var clouds []Cloud
	result, err := db.Query("SELECT * FROM clouds")
	checkErr(err)
	defer result.Close()
	for result.Next() {
		var cloud Cloud
		err = result.Scan(&cloud.CloudID, &cloud.CloudTrainingPath, &cloud.CloudTrainingAvailable, &cloud.CloudPointOfContact, &cloud.ExternalTrainingPoc)
		checkErr(err)
		clouds = append(clouds, cloud)
	}
	json.NewEncoder(w).Encode(clouds)
}

// Get Single CloudID
func GetCloudID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result, err := db.Query("SELECT * FROM clouds WHERE cloud_id = $1", params["cloudid"])
	checkErr(err)
	defer result.Close()
	var cloud Cloud
	for result.Next() {
		err := result.Scan(&cloud.CloudID, &cloud.CloudTrainingPath, &cloud.CloudTrainingAvailable, &cloud.CloudPointOfContact, &cloud.ExternalTrainingPoc)
		checkErr(err)
	}
	json.NewEncoder(w).Encode(cloud)
}

// Create CloudID
func CreateCloudID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	stmt, err := db.Prepare("INSERT INTO clouds(cloud_id, cloudtraining_path, cloudtraining_available, cloud_point_of_contact, external_training_poc) VALUES ($1, $2 , $3, $4, $5, $6)")
	checkErr(err)
	body, err := ioutil.ReadAll(r.Body)
	checkErr(err)
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	cloud_id := keyVal["cloudid"]
	cloudtraining_path := keyVal["cloudtrainingpath"]
	cloudtraining_available := keyVal["cloudtrainingavailable"]
	cloud_point_of_contact := keyVal["cloudpointofcontact"]
	external_training_poc := keyVal["externaltrainingpoc"]
	_, err = stmt.Exec(cloud_id, cloudtraining_path, cloudtraining_available, cloud_point_of_contact, external_training_poc)
	checkErr(err)
	fmt.Fprintln(w, "New CloudID Has Been Created")
}

// Update CloudID
func UpdateCloudID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	stmt, err := db.Prepare("UPDATE clouds SET cloudtraining_path = $1, cloudtraining_available = $2, cloud_point_of_contact = $3, external_training_poc = $4, WHERE cloud_id = $6")
	checkErr(err)
	body, err := ioutil.ReadAll(r.Body)
	checkErr(err)
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	cloudtraining_path := keyVal["cloudtrainingpath"]
	cloudtraining_available := keyVal["cloudtrainingavailable"]
	cloud_point_of_contact := keyVal["cloudpointofcontact"]
	external_training_poc := keyVal["externaltrainingpoc"]
	_, err = stmt.Exec(cloudtraining_path, cloudtraining_available, cloud_point_of_contact, external_training_poc, params["cloudid"])
	checkErr(err)
	fmt.Fprintf(w, "CloudID with ID = %s was Updated", params["cloudid"])
}

// Delete CloudID
func DeleteCloudID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	stmt, err := db.Prepare("DELETE FROM clouds WHERE cloud_id = $1")
	checkErr(err)
	_, err = stmt.Exec(params["cloudid"])
	checkErr(err)
	fmt.Fprintf(w, "CloudID with ID = %s Has Been Deleted", params["cloudid"])
}

// Credentials Methods

// Get All Credentials
func GetCredentials(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var credentials []Credentials
	result, err := db.Query("SELECT * FROM credentials")
	checkErr(err)
	defer result.Close()
	for result.Next() {
		var credential Credentials
		err = result.Scan(&credential.Email, &credential.Password, &credential.Role)
		checkErr(err)
		credentials = append(credentials, credential)
	}
	json.NewEncoder(w).Encode(credentials)
}

// Get Single Credential
func GetCredential(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result, err := db.Query("SELECT * FROM credentials WHERE email=$1", params["email"])
	checkErr(err)
	defer result.Close()
	var credential Credentials
	for result.Next() {
		err := result.Scan(&credential.Email, &credential.Password, &credential.Role)
		checkErr(err)
	}
	json.NewEncoder(w).Encode(credential)
}

// Create Credentials
func CreateCredentials(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	stmt, err := db.Prepare("INSERT INTO credentials(email, password, role) VALUES ($1, $2, $3)")
	checkErr(err)
	body, err := ioutil.ReadAll(r.Body)
	checkErr(err)
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	email := keyVal["email"]
	password := keyVal["password"]
	role := keyVal["role"]
	_, err = stmt.Exec(email, password, role)
	checkErr(err)
	fmt.Fprintln(w, "New Credentials Have Been Created")
}

// Update Credentials
func UpdateCredentials(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	stmt, err := db.Prepare("UPDATE credentials SET email=$1, password=$2, role=$3 WHERE email = $4")
	checkErr(err)
	body, err := ioutil.ReadAll(r.Body)
	checkErr(err)
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	email := keyVal["email"]
	password := keyVal["password"]
	role := keyVal["role"]
	_, err = stmt.Exec(email, password, role, params["email"])
	checkErr(err)
	fmt.Fprintf(w, "Credentials with email = %s was Updated", params["email"])
}

// Delete Credentials
func DeleteCredentials(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	stmt, err := db.Prepare("DELETE FROM credentials WHERE email = $1")
	checkErr(err)
	_, err = stmt.Exec(params["email"])
	checkErr(err)
	fmt.Fprintf(w, "Credentials with email = %s Has Been Deleted", params["email"])
}

//SignIn
func SignIn(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	checkErr(err)

	var user Credentials
	result, err := db.Query("SELECT email, password, role FROM credentials WHERE email=$1", creds.Email)

	for result.Next(){
		err:= result.Scan(&user.Email, &user.Password, &user.Role)
		checkErr(err)
	}
	if user.Email == ""{
		err := errors.New("Email or Password is incorrect")
		json.NewEncoder(w).Encode(err)
		fmt.Println("Email or password is incorrect")
		return
	}
	check := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password))
	fmt.Println(check)
	if check != nil {
		err := errors.New("Email or Password is incorrect")
		json.NewEncoder(w).Encode(err)
		return
	}
	if creds.Role != user.Role{
		err:= errors.New("Role is incorrect")
		json.NewEncoder(w).Encode(err)
		return
	}
	validToken, err := GenerateJWT(user.Email, user.Role)
	checkErr(err)
	var token Token
	token.Email = user.Email
	token.Role = user.Role
	token.TokenString = validToken
	json.NewEncoder(w).Encode(token)
}

//SignUp
func SignUp(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var user Credentials
	err := json.NewDecoder(r.Body).Decode(&user)
	checkErr(err)
	
	var dbuser Employee
	result, err := db.Query("SELECT email FROM employees WHERE email=$1", user.Email)
	checkErr(err)
	for result.Next(){
		err := result.Scan(&dbuser.Email)
		checkErr(err)
	}
	if dbuser.Email != "" {
		err := errors.New("Email already in use")
		json.NewEncoder(w).Encode(err)
		return
	}
	user.Password = getHash([]byte(user.Password))

	stmt, err := db.Prepare("INSERT INTO credentials(email, password, role) VALUES ($1, $2, $3)")
	_, err = stmt.Exec(user.Email, user.Password, user.Role)
	checkErr(err)
	json.NewEncoder(w).Encode(user)
}

// MiddleWare Function
func IsAuthorized(handler http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){

		if r.Header["Token"] == nil {
			err := errors.New("No Token Found")
			json.NewEncoder(w).Encode(err)
			return
		}

		var SECRET_KEY = []byte("gosecretkey")

		token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error){
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an errror in parsing")
			}	
			return SECRET_KEY, nil
		})

		if err != nil {
			err := errors.New("Your Token Has Expired")
			json.NewEncoder(w).Encode(err)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid{
			if claims["role"] == "admin" {
				r.Header.Set("Role", "admin")
				handler.ServeHTTP(w, r)
				return
			}else if claims["role"] == "user" {
				r.Header.Set("Role", "user")
				handler.ServeHTTP(w, r)
				return
			}
		}
		reserr := errors.New("Not Authorized")
		json.NewEncoder(w).Encode(reserr)
	}
}

func AdminIndex(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Role") != "admin" {
		w.Write([]byte("Not authorized."))
		return
	}
	w.Write([]byte("Welcome, Admin."))
}

func UserIndex(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Role") != "user" {
		w.Write([]byte("Not Authorized."))
		return
	}
	w.Write([]byte("Welcome, User."))
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
            "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
    }
    // Stop here if its Preflighted OPTIONS request
    if req.Method == "OPTIONS" {
        return
    }

    c.R.ServeHTTP(rw, req)
}