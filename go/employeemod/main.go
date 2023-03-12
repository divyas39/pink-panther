package main

import (
	"fmt"
	// "net/http"
	"controller"

	"github.com/gorilla/mux"
)

func addApproutes() {

	route := mux.NewRouter()

	route.HandleFunc("/employees", controller.getEmployees).Methods("GET")

	route.HandleFunc("/employee", controller.insertEmployees).Methods("POST")

	route.HandleFunc("/employee/{id}", controller.deleteEmployee).Methods("DELETE")

	route.HandleFunc("/employee", controller.updateEmployee).Methods("PUT")

	fmt.Println("Routes are Loded.")
}

func main() {
	addApproutes()
}