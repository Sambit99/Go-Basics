package main

import "fmt"

func main() {
    fmt.Println("Class on pointers")

	var ptr *int;
	fmt.Println("Value of pointert is ",ptr)

	mynumber := 23

	var ptrnum = &mynumber

	fmt.Println("Value of pointert is ",ptrnum) 
	fmt.Println("Value of pointert is ",*ptrnum) 

	*ptrnum = *ptrnum * 2

	fmt.Println("New value is ",mynumber)

}