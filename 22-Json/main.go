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
	DecodeJson()
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

func DecodeJson() {
	jsonData := []byte(`
        {
			"courseName": "ReactJS",
			"Price": 299,
			"website": "in",
			"tags": ["web-dev","js"]
        }`)

	var course course

	checkValid := json.Valid(jsonData)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonData, &course)
		fmt.Printf("%#v\n", course)
	} else {
		fmt.Println("JSON is not valid")
	}

	// some cases when we want to add key value pair

	var myData map[string]interface{}
	json.Unmarshal(jsonData, &myData)
	fmt.Printf("%#v\n", myData)

	for key, val := range myData {
		fmt.Printf("Key : %v\nValue:%v\n", key, val)
	}
}
