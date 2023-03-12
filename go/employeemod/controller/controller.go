package controller

import (
	"encoding/json"
	"net/http"
	"../service"
	"../model"
	"github.com/gorilla/mux"
)

func getEmployees(response http.ResponseWriter, request *http.Request) {
	var httpError = model.ErrorResponse{
		Code: http.StatusInternalServerError, Message: "No employees in database",
	}
	jsonResponse := service.GetEmployeesFromDB()

	if jsonResponse == nil {
		returnErrorResponse(response, request, httpError)
	} else {
		response.Header().Set("Content-Type", "application/json")
		response.Write(jsonResponse)
	}
}

func insertEmployees(response http.ResponseWriter, request *http.Request) {
	var httpError = model.ErrorResponse{
		Code: http.StatusInternalServerError, Message: "Invalid input",
	}
	var employeeDetails model.Employee
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&employeeDetails)
	defer request.Body.Close()
	if err != nil {
		returnErrorResponse(response, request, httpError)
	} else {
		httpError.Code = http.StatusBadRequest
		if employeeDetails.Name == "" {
			httpError.Message = "Name can't be empty"
			returnErrorResponse(response, request, httpError)
		} else if employeeDetails.Experience < 0 {
			httpError.Message = "Experience can't be a negative number"
			returnErrorResponse(response, request, httpError)
		} else {
			isInserted := service.InsertEmployeeInDB(employeeDetails)
			if isInserted {
				getEmployees(response, request)
			} else {
				returnErrorResponse(response, request, httpError)
			}
		}
	}
}

func deleteEmployee(response http.ResponseWriter, request *http.Request) {
	var httpError = model.ErrorResponse{
		Code: http.StatusInternalServerError, Message: "Employee ID not found",
	}
	userID := mux.Vars(request)["id"]
	if userID == "" {
		httpError.Message = "Employee ID can't be empty"
		returnErrorResponse(response, request, httpError)
	} else {
		isdeleted := service.DeleteEmployeeFromDB(userID)
		if isdeleted {
			getEmployees(response, request)
		} else {
			returnErrorResponse(response, request, httpError)
		}
	}
}

func updateEmployee(response http.ResponseWriter, request *http.Request) {
	var httpError = model.ErrorResponse{
		Code: http.StatusInternalServerError, Message: "It's not you it's me.",
	}
	var employeeDetails model.Employee
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&employeeDetails)
	defer request.Body.Close()
	if err != nil {
		returnErrorResponse(response, request, httpError)
	} else {
		httpError.Code = http.StatusBadRequest
		if employeeDetails.Name == "" {
			httpError.Message = "Name can't be empty"
			returnErrorResponse(response, request, httpError)
		} else if employeeDetails.Id == "" {
			httpError.Message = "Employee Id can't be empty"
			returnErrorResponse(response, request, httpError)
		} else if employeeDetails.Experience < 0 {
			httpError.Message = "Experience can't be negative"
			returnErrorResponse(response, request, httpError)
		} else {
			isUpdated := service.UpdateEmployeeInDB(employeeDetails)
			if isUpdated {
				getEmployees(response, request)
			} else {
				returnErrorResponse(response, request, httpError)
			}
		}
	}
}

func returnErrorResponse(response http.ResponseWriter, request *http.Request, errorMesage model.ErrorResponse) {
	httpResponse := &model.ErrorResponse{Code: errorMesage.Code, Message: errorMesage.Message}
	jsonResponse, err := json.Marshal(httpResponse)
	if err != nil {
		panic(err)
	}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(errorMesage.Code)
	response.Write(jsonResponse)
}
