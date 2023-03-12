package service

import (
	// "context"
	"../model"
	// "../service"
	"../config"
	// "database/sql"
	// "errors"
	"fmt"
	"encoding/json"
)

func GetEmployeesFromDB() []byte {
	var (
		employee  model.Employee
		employees []model.Employee
	)

	var mysql = config.Connect_sql()

	rows, err := mysql.Query("SELECT * FROM employees")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for rows.Next() {
		rows.Scan(&employee.Id, &employee.Name, &employee.Experience)
		employees = append(employees, employee)
	}
	defer rows.Close()
	jsonResponse, jsonError := json.Marshal(employees)
	if jsonError != nil {
		fmt.Println(jsonError)
		return nil
	}
	return jsonResponse
}

func InsertEmployeeInDB(employeeDetails model.Employee) bool {
	var mysql = config.Connect_sql()
	stmt, err := mysql.Prepare("INSERT into employees SET name=?,experience=?")
	if err != nil {
		fmt.Println(err)
		return false
	}
	_, queryError := stmt.Exec(employeeDetails.Name, employeeDetails.Experience)
	if queryError != nil {
		fmt.Println(queryError)
		return false
	}
	return true
}

func DeleteEmployeeFromDB(employeeID string) bool {
	var mysql = config.Connect_sql()
	stmt, err := mysql.Prepare("DELETE FROM employees WHERE id=?")
	if err != nil {
		fmt.Println(err)
		return false
	}
	_, queryError := stmt.Exec(employeeID)
	if queryError != nil {
		fmt.Println(queryError)
		return false
	}
	return true
}

func UpdateEmployeeInDB(employeeDetails model.Employee) bool {
	var mysql = config.Connect_sql()
	stmt, err := mysql.Prepare("UPDATE employees SET name=?,experience=? WHERE id=?")
	if err != nil {
		fmt.Println(err)
		return false
	}
	_, queryError := stmt.Exec(employeeDetails.Name, employeeDetails.Experience, employeeDetails.Id)
	if queryError != nil {
		fmt.Println(queryError)
		return false
	}
	return true
}