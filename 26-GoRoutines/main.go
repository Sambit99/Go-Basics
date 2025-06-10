package main

import (
	"fmt"
	"time"
)

func main() {
	go greeter("Hello")
	greeter("World")

}

func greeter(name string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(name)
	}
}
