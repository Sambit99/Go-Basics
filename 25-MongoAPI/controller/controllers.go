package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Sambit99/Go-Basics/MongoAPI/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const connectionString = "mongodb://localhost:27017"
const dbName = "netflix"
const colName = "watchlist"

var collection *mongo.Collection

var ctx = context.Background()

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

// Update 1 record
func updateOneMovie(movieId string) {
	id, err := bson.ObjectIDFromHex(movieId)

	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Modified count: ", result.ModifiedCount)
}

// Delete 1 record
func deleteOneMovie(movieId string) {
	id, err := bson.ObjectIDFromHex(movieId)

	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}

	result, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted Record count: ", result.DeletedCount)
}

// Delete all movies
func deleteAllMovies() int64 {

	result, err := collection.DeleteMany(context.Background(), bson.M{}, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted Records count: ", result.DeletedCount)

	return result.DeletedCount
}

// Get all movies
func getAllMovies() []bson.M {
	cursor, err := collection.Find(ctx, bson.M{}, nil)

	if err != nil {
		log.Fatal(err)
	}

	var movies []bson.M
	for cursor.Next(ctx) {
		var movie bson.M

		if err = cursor.Decode(&movie); err != nil {
			log.Fatal(err)
		}

		movies = append(movies, movie)
	}

	defer cursor.Close(ctx)

	fmt.Println("Records: ", movies)

	return movies
}

// Actual Controllers

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")

	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}
