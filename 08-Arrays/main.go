package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays in Golang")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Orange"
	fruitList[3] = "Guava"

	fmt.Println("Fruit List is : ",fruitList) 

	fmt.Println("Fruit List is : ",len(fruitList))  

	var vegList = [3]string{"Potato","Onion","Tomato"}

	fmt.Println("Veg List is : ",vegList)
	fmt.Println("Veg List is : ",len(vegList))

}