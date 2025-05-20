package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("File in Golang")

	content := "This is written by code in files"

	file,err := os.Create("./myTextfile.txt")

	if err != nil {
		panic(err)
	}

	length,err := io.WriteString(file,content)

	checkNilErr(err)

	fmt.Println("Length is : ",length)
	defer file.Close()
	readFile("./myTextfile.txt")
}

func readFile(filename string) {
	databyte,err := os.ReadFile(filename)
	checkNilErr(err)

	fmt.Println("Text data inside the file is \n",string(databyte))
}
func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}