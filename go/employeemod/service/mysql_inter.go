package service

import (
	"employeemod/model"
)

type MysqlEmployee interface {
	InsertEmployeeInDB(employeeDetails model.Employee) (bool, error)

	GetEmployeebyID(employeeId string) (model.Employee, error)

	GetEmployeesFromDB() ([]byte, error)

	UpdateEmployeeInDB(employeeDetails model.Employee) (bool, error)

	DeleteEmployeeFromDB(employeeID string) (bool, error)
}
