package api

import (
	"log"
	"net/http"
	handle"github.com/homirock/GoProjects/Handler"
	"github.com/gorilla/mux"
)

func ApiRequest() {
	route := mux.NewRouter()
	route.HandleFunc("/emp/{id}", handle.Homepage).Methods("GET").Host("localhost").Schemes("http")
	// route.HandleFunc("/employees", getAllEmployees)
	// route.HandleFunc("/employee", createEmployee).Methods("POST")
	// route.HandleFunc("/employee/{id}", deleteEmployee).Methods("DELETE")
	// route.HandleFunc("/employee/{id}", getEmployee)
	log.Fatal(http.ListenAndServe(":4000", route))
}
