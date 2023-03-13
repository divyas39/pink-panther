package config

// import (
// 	"context"
// 	"fmt"
// 	"log"

// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	EmployeesCollection     *mongo.Collection
	Ctx                 = context.TODO()
)


// func ConnectDB() *mongo.Collection {

// 	// Set client options
// 	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017")

// 	// Connect to MongoDB
// 	client, err := mongo.Connect(context.TODO(), clientOptions)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Connected to MongoDB!")

// 	collection := client.Database("employees").Collection("employees")

// 	return collection
// }


/*Setup opens a database connection to mongodb*/
func Setup() *mongo.Collection {
	host := "127.0.0.1"
	port := "27017"
	connectionURI := "mongodb://" + host + ":" + port + "/"
	clientOptions := options.Client().ApplyURI(connectionURI)
	client, err := mongo.Connect(Ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(Ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("employees")
	EmployeesCollection = db.Collection("employees")
	fmt.Println("connected to database")
	return EmployeesCollection
}