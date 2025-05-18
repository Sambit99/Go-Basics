package main

import "fmt"

func main() {
	fmt.Println("Maps in GoLang")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"


	fmt.Println("List of all languages : ",languages) 
	fmt.Println("JS Shorts for : ",languages["JS"]) 

	delete(languages,"RB")
	fmt.Println("List of all languages : ",languages) 


	// loops 

	for key,value := range languages{
		fmt.Printf("For Key %v , value is %v\n",key,value) 
	}
}