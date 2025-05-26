package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"courseName"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("JSON")
	EncodeJson()
}

func EncodeJson() {
	myCourse := []course{
		{"ReactJS", 299, "in", "Kola", []string{"web-dev", "js"}},
		{"AngularJS", 199, "in", "Kola", []string{"web-dev", "js"}},
		{"NextJS", 599, "in", "Kola", nil},
		{"NestJS", 399, "in", "Kola", []string{"web-dev", "js"}},
	}

	// package the data a JSON Data

	finalJson, err := json.MarshalIndent(myCourse, "", "\t")

	if err != nil {
		panic(err)
	}

	// fmt.Println(myCourse)
	fmt.Printf("%s\n", finalJson)
}
