package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to Pizza App")
	fmt.Println("Please rate our pizza between 1 and 5")

	reader := bufio.NewReader(os.Stdin)

	input,_ := reader.ReadString('\n')
	fmt.Println("Thanks for rating",input)

	numRating,err := strconv.ParseFloat(strings.TrimSpace(input),64)

	if(err != nil) {
		panic(err)
	}else{
		fmt.Println("Adding 1 to your rating: ",numRating+1)
	}
}