package main

import (
	"fmt"
	"net/url"
)

const myUrl = "https://www.youtube.com/watch?v=akJYqeKDRho"

func main() {
	fmt.Println("Welcome to handling URLs in Go")
	fmt.Println(myUrl)

	// Parsing
	result,_ := url.Parse(myUrl)

	fmt.Println(result) 
	fmt.Println(result.Scheme) 
	fmt.Println(result.Host) 
	fmt.Println(result.Path) 
	fmt.Println(result.Port()) 
	fmt.Println(result.RawQuery) 

	queryParam := result.Query()
	fmt.Printf("The type of query param are : %T\n",queryParam) 
	fmt.Println(queryParam)

	for key,value := range(queryParam) {
		fmt.Println(key,value)
	}

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "www.youtube.com",
		Path: "/watch",
		RawQuery: "v=akJYqeKDRho",
	}

	anotherUrl := partsOfUrl.String()

	fmt.Println(anotherUrl)
}