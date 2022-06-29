package main

import (
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Book struct (Model)
type Employee struct {
	ID          string `json:"id"`
	EmpName     string `json:"empname"`
	Dateofbirth string `json:"dateofbirth"`
	Gender      string `json:"gender"`
	City        string `json:"city"`
	State       string `json:"state"`
	Email       string `json:"email"`
	Phonenumber string `json:"phonenumber"`
}

//Init employee Var as slice Employee Struct
//var employees []Employee

//var comn = "server=localhost\\SQLEXPRESS;user id=xyz;password=yourpassword;port=1433;database=yourdatabasename;"
var comn = "server=localhost;user id=pari;password=pari1234;port=1433;database=Emplyoee;"

//Get Employees
func getEmployees(w http.ResponseWriter, r *http.Request) {
	db, err := sqlx.Open("mssql", comn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//employees := Employee{}
	tsql := fmt.Sprintf(`[dbo].[GetEmplyoeeDetail]`)
	rows, err := db.Queryx(tsql)

	if err != nil {
		log.Fatal(err)
	}
	var employees []Employee
	for rows.Next() {
		var id string
		var empname string
		var dateofbirth string
		var gender string
		var city string
		var state string
		var email string
		var phonenumber string

		err2 := rows.Scan(&id, &empname, &dateofbirth, &gender, &city, &state, &email, &phonenumber)

		if err2 != nil {
			log.Fatal(err2)
		}
		js, err2 := json.Marshal(err2)
		e(err2)
		employee := Employee{id, empname, dateofbirth, gender, city, state, email, phonenumber}
		employees = append(employees, employee)
		fmt.Println(string(js))
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

//Get Single employee
func getEmployee(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	db, err := sqlx.Open("mssql", comn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	//var employee *Employee = &Employee{}
	employee := Employee{}
	employee.ID = params["id"]
	//_ = json.NewDecoder(r.Body).Decode(&employee)
	tsql := fmt.Sprintf(` [dbo].[GetSingleEmployeeDetail]'%s'`, employee.ID)

	rows, err := db.Queryx(tsql)

	if err != nil {
		log.Fatal(err)
	}
	var employees []Employee
	for rows.Next() {
		var id string
		var empname string
		var dateofbirth string
		var gender string
		var city string
		var state string
		var email string
		var phonenumber string

		err2 := rows.Scan(&id, &empname, &dateofbirth, &gender, &city, &state, &email, &phonenumber)
		if err2 != nil {
			log.Fatal(err2)
		}
		js, err2 := json.Marshal(err2)
		e(err2)
		employee := Employee{id, empname, dateofbirth, gender, city, state, email, phonenumber}
		employees = append(employees, employee)
		fmt.Println(string(js))
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)

}

// fmt.Print("Select Single Employee")
// w.Header().Set("Server", "Select Single Employee")
// w.WriteHeader(200)
// defer stmit.Close()

//Create Employee
func createEmployees(w http.ResponseWriter, r *http.Request) {
	db, err := sqlx.Open("mssql", comn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	var employee Employee
	json.NewDecoder(r.Body).Decode(&employee)
	fmt.Println(employee)
	tsql := fmt.Sprintf(`[dbo].[InsertEmp] '%s','%s','%s','%s','%s','%s','%s'  `, employee.EmpName, employee.Dateofbirth, employee.Gender, employee.City, employee.State, employee.Email, employee.Phonenumber)
	stmit, err := db.Queryx(tsql)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Employee was Created")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Server", "Employee was Created")
	w.WriteHeader(200)
	defer stmit.Close()

}

//Update Employee
func updateEmployee(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	db, err := sqlx.Open("mssql", comn)
	e(err)
	defer db.Close()
	employee := Employee{}
	employee.ID = params["id"]

	_ = json.NewDecoder(r.Body).Decode(&employee)

	tsql := fmt.Sprintf(` [dbo].[updateEmployeeDetail] '%s','%s','%s','%s','%s','%s','%s','%s' `, employee.ID, employee.EmpName, employee.Dateofbirth, employee.Gender, employee.City, employee.State, employee.Email, employee.Phonenumber)

	updatestm, err := db.Queryx(tsql)
	if err != nil {
		e(err)
	}
	// lastAffec, err := updatestm.LastInsertId()
	// if err != nil {
	// log.Fatal(err)
	// }
	// fmt.Print(lastAffec)

	// fmt.Print("Employee was Update")
	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(employees)
	// w.Header().Set("Server", "Employee was Updated")
	// w.WriteHeader(200)
	defer updatestm.Close()
}

//Delete Employee
func deleteEmployee(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	db, err := sqlx.Open("mssql", comn)
	e(err)
	defer db.Close()
	var employee Employee
	employee.ID = params["id"]
	_ = json.NewDecoder(r.Body).Decode(&employee)
	tsql := fmt.Sprintf(`[dbo].[DeleteEmployeeDetail]'%s'`, employee.ID)
	stmit, err := db.Queryx(tsql)

	if err != nil {
		e(err)
	}
	fmt.Print("Employee was deleted")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Server", "Employee was deleted")
	w.WriteHeader(200)
	defer stmit.Close()
}

func main() {
	//Innit Router
	r := mux.NewRouter()

	//Rounter Handler //End Point
	r.HandleFunc("/api/Employees", getEmployees).Methods("GET")
	r.HandleFunc("/api/Employee/{id}", getEmployee).Methods("GET")
	r.HandleFunc("/api/Employees", createEmployees).Methods("POST")
	r.HandleFunc("/api/Employee/{id}", updateEmployee).Methods("PUT")
	r.HandleFunc("/api/Employee/{id}", deleteEmployee).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8087", handlers.CORS(handlers.AllowedMethods([]string{"PUT", "POST", "DELETE", "GET", "HEAD"}), handlers.AllowedHeaders([]string{"authentication", "Content-Type"}), handlers.AllowedOrigins([]string{"*"}), handlers.AllowCredentials())(r)))
}

//Error Checking
func e(err error) {
	if err != nil {
		log.Fatal(err, "connection failed")
	}
}
