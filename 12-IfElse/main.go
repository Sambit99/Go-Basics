package main

import "fmt"

func main() {
	fmt.Println("If Else in Go")

	loginCount := 1

	var result string

	if loginCount < 10 {
		result = "Regular User"
	}else if loginCount > 10 {
		result = "New User"
	}else{
		result = "Exactly 10 login count"
	}

	fmt.Println(result)

	if num := 5; num > 10{
		fmt.Println("Num is greater than 10")
	}else{
		fmt.Println("Num is less than 10")
	}

}