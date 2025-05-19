package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch and case in Golang")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1

	fmt.Println("Value of dice is ",diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("Dice value is 2. You can move to 2 spot.")
	case 3:
		fmt.Println("Dice value is 3. You can move to 3 spot.")
	case 4:
		fmt.Println("Dice value is 4. You can move to 4 spot.")
	fallthrough
	case 5:
		fmt.Println("Dice value is 5. You can move to 5 spot.")
	fallthrough
	case 6:
		fmt.Println("Dice value is 6. You can move to 6 spot and roll the dice again")
	default:
		fmt.Println("LOL")
	}
}