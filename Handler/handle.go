package handle

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	m "github.com/homirock/GoProjects/model"
)

func Homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Homepage")
}

func AllemployeesDetails(w http.ResponseWriter, r *http.Request) {
	emp1 := m.Employee{Id: "1", Name: "partha", MailId: "pgiri@gmail.com"}
	emp2 := m.Employee{Id: "2", Name: "soumya", MailId: "soumya@gmail.com"}
	emps := []m.Employee{emp1, emp2}
	json.NewEncoder(w).Encode(emps)
}
func EmployeesDetails(w http.ResponseWriter, r *http.Request) {
	emp1 := m.Employee{Id: "1", Name: "partha", MailId: "pgiri@gmail.com"}
	emp2 := m.Employee{Id: "2", Name: "soumya", MailId: "soumya@gmail.com"}
	emps := []m.Employee{emp1, emp2}
	vars := mux.Vars(r)
	key := vars["id"]
	for _, employee := range emps {
		if employee.Id == key {
			json.NewEncoder(w).Encode(employee)
		}
	}
}
func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	emp1 := m.Employee{Id: "1", Name: "partha", MailId: "pgiri@gmail.com"}
	emp2 := m.Employee{Id: "2", Name: "soumya", MailId: "soumya@gmail.com"}
	emps := []m.Employee{emp1, emp2}
	reqBody, _ := ioutil.ReadAll(r.Body)
	var employee m.Employee
	json.Unmarshal(reqBody, &employee)
	emps = append(emps, employee)
	json.NewEncoder(w).Encode(employee)
}
