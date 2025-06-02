package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Go Modules")
	greeter()

	r := mux.NewRouter()
	r.HandleFunc("/home", serveHome).Methods("GET")
	r.HandleFunc("/", Req)
	// http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":4000", r))
}

func greeter() {
	fmt.Println("Hello Users")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello World</h1>"))
}

func Req(w http.ResponseWriter, r *http.Request) {
	host := r.Host
	uri := r.RequestURI

	fmt.Println("Host: ", host)
	fmt.Println("URI: ", uri)
}
