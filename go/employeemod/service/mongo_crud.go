package service

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	// "log"
	"employeemod/model"
	"fmt"
	"encoding/json"
	conf "employeemod/config"
)

var (
	EmployeesCollection     *mongo.Collection
	Ctx                 = context.TODO()
)

type MongoEmployeeStruct struct{
	
}

func (o *MongoEmployeeStruct) CreateEmployee(emp model.Employee) (bool, error) {
	EmployeesCollection = conf.Setup()
	fmt.Println("employee", emp)
	result, err := EmployeesCollection.InsertOne(Ctx, emp)
	fmt.Println("ERRRRRR", err)
	if err != nil {
		return false, err
	}
	fmt.Println(result)
	return 	true, nil
}


func (o *MongoEmployeeStruct) GetEmployee(id string) (model.Employee, error) {
	EmployeesCollection = conf.Setup()
	var emp model.Employee
	// objectId, err := primitive.ObjectIDFromHex(id)
	// if err != nil{
	// 	return emp,err
	// }
	
	err := EmployeesCollection.FindOne(Ctx, bson.D{{"_id", id}}).Decode(&emp)
	if err != nil {
		return emp, err
	}
	return emp, nil
}


func (o *MongoEmployeeStruct) GetEmployees() ([]byte, error) {
	EmployeesCollection = conf.Setup()
	var emp model.Employee
	var emps []model.Employee

	cursor, err := EmployeesCollection.Find(Ctx, bson.D{})
	jsonResponse, jsonError := json.Marshal(emps)
	if err != nil {
		defer cursor.Close(Ctx)
		return jsonResponse, jsonError
	}

	for cursor.Next(Ctx) {
		err := cursor.Decode(&emp)
		if err != nil {
			return json.Marshal(emps)
		}
		emps = append(emps, emp)
	}

	return json.Marshal(emps)
}


func (o *MongoEmployeeStruct) UpdateEmployee(employeeDetails model.Employee) (bool, error) {
	EmployeesCollection = conf.Setup()
	id := employeeDetails.Id
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"name", employeeDetails.Name}, {"experiemce", employeeDetails.Experience}}}}
	_, err := EmployeesCollection.UpdateOne(
		Ctx,
		filter,
		update,
	)
	if err != nil{
		return false, err
	}
	return true, nil
}


func (o *MongoEmployeeStruct) DeleteEmployee(id string) (bool, error) {
	EmployeesCollection = conf.Setup()
	_, err := EmployeesCollection.DeleteOne(Ctx, bson.D{{"_id", id}})
	if err != nil {
		return false, err
	}
	return true, nil
}