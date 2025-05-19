package main

import (
	"fmt"
)

func main() {

	fmt.Println("Welcome to functions in Golang")
	greeter()

	result := addition(3,5)

	fmt.Println("Result is: ",result)

	result,message := addMultiple(1,2,3,4,5,6,7,8,9)

	fmt.Println("Result is: ",result)
	fmt.Println("Message is: ",message)
}

func addition(x int, y int) int { 
	return x+y;
}

func addMultiple(values ...int) (int,string) {
	total := 0
	for _,val := range values{
	total += val
	}
	return total,"The function executed successfully"
}

func greeter(){
	fmt.Println("Namaste")
}