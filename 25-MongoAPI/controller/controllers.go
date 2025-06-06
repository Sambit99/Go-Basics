package controller

import (
	"context"
	"fmt"
	"log"

	"github.com/Sambit99/Go-Basics/MongoAPI/model"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const connectionString = "mongodb://localhost:27017"
const dbName = "netflix"
const colName = "watchlist"

var collection *mongo.Collection

func init() {
	// Mongo Client Options
	clientOptions := options.Client().ApplyURI(connectionString)

	// Mongo Client Connection
	client, err := mongo.Connect(clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Ping the database to verify connection
	if err := client.Ping(context.TODO(), nil); err != nil {
		log.Fatal("Could not connect to MongoDB: " + err.Error())
	}

	fmt.Println("Mongodb Connection Successful")

	collection = client.Database(dbName).Collection(colName)

	fmt.Println("Collection instance is ready")

	// defer func() {
	// 	if err = client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()
}

// MongoDB Helpers

// insert 1 record
func insertOneMovie(movie model.Netflix) {
	newRecord, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted 1 movie in db with id: ", newRecord.InsertedID)
}
