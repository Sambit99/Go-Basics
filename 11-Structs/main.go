package main

import "fmt"

func main() {
	fmt.Println("Welcome to Structs");

	sambit := User{"Sambit","Sambit.sahoo@go.dev",true,25}
	fmt.Println(sambit)
 
	fmt.Printf("User details %+v\n",sambit)
}

type User struct{
	Name string
	Email string
	Status bool
	Age int
}