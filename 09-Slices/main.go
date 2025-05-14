package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices");

	var fruitList = []string{"Apple","Banana","Orange"}
 
	fmt.Printf("Type of fruitlist %T\n",fruitList)

	fruitList = append(fruitList, "Mango","Guava") 
	fmt.Println(fruitList) 
	
	fruitList = append(fruitList[1:3])

	fmt.Println(fruitList) 

	highScore := make([]int,4)
	highScore[0] = 123;
	highScore[1] = 323;
	highScore[2] = 134;
	highScore[3] = 232;
	// highScore[4] = 232; 

	highScore =append(highScore, 255,143,343) 

	fmt.Println(highScore) 

	sort.Ints(highScore)
	fmt.Println(highScore) 

	// Remove a value in Slices based on Index

	var courses = []string{"ReactJs","Javascript","Swift","Python","GoLang","Ruby"}
	fmt.Println(courses) 
	
	var index int = 2;

	courses = append(courses[:index],courses[index+1:]...)
	fmt.Println(courses)

}