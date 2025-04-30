package main

import (
	"fmt"
	"math/rand"
	"crypto/rand"
	"math/big"
)

func main() {
	fmt.Println("Welcome to Maths in Go")

	var myNumberOne int = 2
	var myNumberTwo float64 = 4

	fmt.Println("The Sum is ",myNumberOne+int(myNumberTwo))

	// Random Number (Math Library)

	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(5)+1)

	
	// Random Number (Crypto Library)

	myRandomNum,_ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println("My Random Number is :",myRandomNum)


}