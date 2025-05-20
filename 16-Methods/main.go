package main

import "fmt"

func main() {
	fmt.Println("Welcome to Structs")

	sambit := User{"Sambit", "Sambit.sahoo@go.dev", true, 25}
	fmt.Println(sambit) 
	fmt.Printf("User details %+v\n", sambit)
	sambit.GetStatus()
	sambit.NewMail()
	fmt.Printf("Name is %v and Email is %v\n", sambit.Name, sambit.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus(){
	fmt.Println("Is user active: ",u.Status)
}

func (u User) NewMail(){
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ",u.Email)
}
