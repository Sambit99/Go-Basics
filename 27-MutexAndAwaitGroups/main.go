package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race Condition")
	wg := &sync.WaitGroup{}

	var score = []int{0}

	wg.Add(3)
	go func(wg *sync.WaitGroup) {
		fmt.Println("Go Routine No.1")
		score = append(score, 1)
		wg.Done()
	}(wg)

	// wg.Add(1)
	go func(wg *sync.WaitGroup) {
		fmt.Println("Go Routine No.2")
		score = append(score, 2)
		wg.Done()
	}(wg)

	go func(wg *sync.WaitGroup) {
		fmt.Println("Go Routine No.3")
		score = append(score, 3)
		wg.Done()
	}(wg)

	wg.Wait()

	fmt.Println("Score is : ", score)
}
