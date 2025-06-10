package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var wg sync.WaitGroup
var mut sync.Mutex

var signals = []string{"test"}

func main() {
	// go greeter("Hello")
	// greeter("World")

	websites := []string{
		"https://jsonplaceholder.typicode.com/todos/1",
		"https://jsonplaceholder.typicode.com/todos/2",
		"https://google.com",
		"https://github.com",
		"https://youtube.com",
	}

	for _, url := range websites {
		go getStatusCode(url)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println(signals)

}

// func greeter(name string) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(1 * time.Second)
// 		fmt.Println(name)
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done()
	result, err := http.Get(endpoint)

	if err != nil {
		log.Fatal(err)
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d statuscode for %s \n", result.StatusCode, endpoint)

	}

}
