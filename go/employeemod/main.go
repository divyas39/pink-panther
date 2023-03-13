package main

import (
	"fmt"
	"net/http"
	// "net/http"
	"employeemod/controller"
	"log"

	"github.com/gorilla/mux"
)

func addApproutes() {

	route := mux.NewRouter()

	route.HandleFunc("/employees", controller.GetEmployees).Methods("GET")

	route.HandleFunc("/employee", controller.InsertEmployees).Methods("POST")

	route.HandleFunc("/employee/{id}", controller.DeleteEmployee).Methods("DELETE")

	route.HandleFunc("/employee", controller.UpdateEmployee).Methods("PUT")

	fmt.Println("Routes are Loaded.")
	log.Fatal(http.ListenAndServe(":10000", route))
}

func main() {
	addApproutes()
}