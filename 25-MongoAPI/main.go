package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Sambit99/Go-Basics/MongoAPI/router"
)

func main() {
	fmt.Println("MongoDB API")
	r := router.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening on port : 4000")

}
