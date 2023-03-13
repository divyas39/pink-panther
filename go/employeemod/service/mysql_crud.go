package service

import (
	// "context"
	"employeemod/model"
	// "../service"
	"employeemod/config"
	// "database/sql"
	// "errors"
	"fmt"
	"encoding/json"
	// "employeemod/interfaces"
)

type MysqlEmployeeStruct struct{

}

func (s *MysqlEmployeeStruct) GetEmployeesFromDB() ([]byte, error) {
	var (
		employee  model.Employee
		employees []model.Employee
	)

	var mysql = config.Connect_sql()

	rows, err := mysql.Query("SELECT * FROM employees")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	for rows.Next() {
		rows.Scan(&employee.Id, &employee.Name, &employee.Experience)
		employees = append(employees, employee)
	}
	defer rows.Close()
	jsonResponse, jsonError := json.Marshal(employees)
	if jsonError != nil {
		fmt.Println(jsonError)
		return nil, jsonError
	}
	return jsonResponse, jsonError
}

func (s *MysqlEmployeeStruct) InsertEmployeeInDB(employeeDetails model.Employee) (bool, error) {
	var mysql = config.Connect_sql()
	stmt, err := mysql.Prepare("INSERT into employees SET name=?,experience=?")
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	_, queryError := stmt.Exec(employeeDetails.Name, employeeDetails.Experience)
	if queryError != nil {
		fmt.Println(queryError)
		return false, queryError
	}
	return true, nil
}

func (s *MysqlEmployeeStruct) DeleteEmployeeFromDB(employeeID string) (bool, error) {
	var mysql = config.Connect_sql()
	stmt, err := mysql.Prepare("DELETE FROM employees WHERE id=?")
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	_, queryError := stmt.Exec(employeeID)
	if queryError != nil {
		fmt.Println(queryError)
		return false, queryError
	}
	return true, nil
}

func (s *MysqlEmployeeStruct) UpdateEmployeeInDB(employeeDetails model.Employee) (bool, error) {
	var mysql = config.Connect_sql()
	stmt, err := mysql.Prepare("UPDATE employees SET name=?,experience=? WHERE id=?")
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	_, queryError := stmt.Exec(employeeDetails.Name, employeeDetails.Experience, employeeDetails.Id)
	if queryError != nil {
		fmt.Println(queryError)
		return false, queryError
	}
	return true, nil
}