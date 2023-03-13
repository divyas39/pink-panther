package controller

import (
	"encoding/json"
	// "errors"
	"fmt"
	"net/http"

	// "employeemod/service"
	"employeemod/model"
	o "employeemod/service"

	"github.com/gorilla/mux"
)


func GetEmployees(response http.ResponseWriter, request *http.Request) {

	db_name := request.URL.Query().Get("db")
	var jsonResponse []byte 
	var jsonError error
	
	var httpError = model.ErrorResponse{
		Code: http.StatusInternalServerError, Message: "No employees in database",
	}

	if db_name == "mysql" {
		mysqlinter := &o.MysqlEmployeeStruct{}
		jsonResponse, jsonError = mysqlinter.GetEmployeesFromDB()
	}else if db_name == "mongodb" {
		mongointer := &o.MongoEmployeeStruct{}
		jsonResponse, jsonError = mongointer.GetEmployees()
		fmt.Println("entered")
	}else{
		jsonResponse = nil
	}
	fmt.Println("RESPONSE", jsonResponse, jsonError)
	

	if jsonResponse == nil {
		ReturnErrorResponse(response, request, httpError)
	} else {
		fmt.Println("TADA")
		response.Header().Set("Content-Type", "application/json")
		response.Write(jsonResponse)
	}
}


func InsertEmployees(response http.ResponseWriter, request *http.Request) {

	db_name := request.URL.Query().Get("db")
	var isInserted bool
	
	var httpError = model.ErrorResponse{
		Code: http.StatusInternalServerError, Message: "Invalid input",
	}
	var employeeDetails model.Employee
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&employeeDetails)
	defer request.Body.Close()
	if err != nil {
		ReturnErrorResponse(response, request, httpError)
	} else {
		httpError.Code = http.StatusBadRequest
		if employeeDetails.Name == "" {
			httpError.Message = "Name can't be empty"
			ReturnErrorResponse(response, request, httpError)
		} else if employeeDetails.Experience < 0 {
			httpError.Message = "Experience can't be a negative number"
			ReturnErrorResponse(response, request, httpError)
		} else {

			if db_name == "mysql"{
				mysqlinter := &o.MysqlEmployeeStruct{}
				isInserted, err = mysqlinter.InsertEmployeeInDB(employeeDetails)
			}else if db_name == "mongodb"{
				mongointer := &o.MongoEmployeeStruct{}
				isInserted, err = mongointer.CreateEmployee(employeeDetails)
				fmt.Println("Document created")
			}else{
				isInserted = false
			}
			fmt.Println("ERROR", err)
			fmt.Println(isInserted)
			
			// if isInserted {
			// 	GetEmployees(response, request)
			// } else {
			// 	ReturnErrorResponse(response, request, httpError)
			// }
		}
	}
}


func DeleteEmployee(response http.ResponseWriter, request *http.Request) {

	db_name := request.URL.Query().Get("db")
	var isDeleted bool
	var err error

	var httpError = model.ErrorResponse{
		Code: http.StatusInternalServerError, Message: "Employee ID not found",
	}
	userID := mux.Vars(request)["id"]
	if userID == "" {
		httpError.Message = "Employee ID can't be empty"
		ReturnErrorResponse(response, request, httpError)
	} else {

		if db_name == "mysql"{
			mysqlinter := &o.MysqlEmployeeStruct{}
			isDeleted, err = mysqlinter.DeleteEmployeeFromDB(userID)
		}else if db_name == "mongodb"{
			mongointer := &o.MongoEmployeeStruct{}
			isDeleted, err = mongointer.DeleteEmployee(userID)
		}else{
			isDeleted = false
		}
		fmt.Println(err)
		
		if isDeleted {
			GetEmployees(response, request)
		} else {
			ReturnErrorResponse(response, request, httpError)
		}
	}
}


func UpdateEmployee(response http.ResponseWriter, request *http.Request) {

	db_name := request.URL.Query().Get("db")
	var isUpdated bool

	var httpError = model.ErrorResponse{
		Code: http.StatusInternalServerError, Message: "It's not you it's me.",
	}
	var employeeDetails model.Employee
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&employeeDetails)
	defer request.Body.Close()
	if err != nil {
		ReturnErrorResponse(response, request, httpError)
	} else {
		httpError.Code = http.StatusBadRequest
		if employeeDetails.Name == "" {
			httpError.Message = "Name can't be empty"
			ReturnErrorResponse(response, request, httpError)
		} else if employeeDetails.Id == "" {
			httpError.Message = "Employee Id can't be empty"
			ReturnErrorResponse(response, request, httpError)
		} else if employeeDetails.Experience < 0 {
			httpError.Message = "Experience can't be negative"
			ReturnErrorResponse(response, request, httpError)
		} else {

			if db_name == "mysql"{
				mysqlinter := &o.MysqlEmployeeStruct{}
				isUpdated, err = mysqlinter.UpdateEmployeeInDB(employeeDetails)
			}else if db_name == "mongodb"{
				mongointer := &o.MongoEmployeeStruct{}
				isUpdated, err = mongointer.UpdateEmployee(employeeDetails)
			}else{
				isUpdated = false
			}
			fmt.Println(err)
			
			if isUpdated {
				GetEmployees(response, request)
			} else {
				ReturnErrorResponse(response, request, httpError)
			}
		}
	}
}


func ReturnErrorResponse(response http.ResponseWriter, request *http.Request, errorMesage model.ErrorResponse) {
	httpResponse := &model.ErrorResponse{Code: errorMesage.Code, Message: errorMesage.Message}
	jsonResponse, err := json.Marshal(httpResponse)
	if err != nil {
		fmt.Println(err)
	}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(errorMesage.Code)
	response.Write(jsonResponse)
}
