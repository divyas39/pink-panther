package service

import (
	"employeemod/model"
)

type MongoEmployee interface {
	CreateEmployee(emp model.Employee) (bool, error)

	GetEmployee(id string) (model.Employee, error)

	GetEmployees() ([]byte, error)

	UpdateEmployee(employeeDetails model.Employee) (bool, error)

	DeleteEmployee(id string) (bool, error)
}