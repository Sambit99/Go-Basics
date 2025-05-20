package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://jsonplaceholder.typicode.com/todos/1"

func main() {
	fmt.Println("Web Requests in Golang")

	response,error := http.Get(url)

	if(error != nil) {
		panic(error)
	}

	fmt.Printf("Response of type %T\n",response) 

	defer response.Body.Close() 

	dataBytes,error := io.ReadAll(response.Body)

	if(error != nil) {
		panic(error)
	}

	content := string(dataBytes)
	fmt.Println("Response ",content)
	fmt.Println("Status Code ",response.StatusCode)
}