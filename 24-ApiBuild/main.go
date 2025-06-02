package main

import "fmt"

// Course Model
type Course struct {
	CourseID   string  `json:"courseid"`
	CourseName string  `json:"coursename"`
	Price      int     `json:"price"`
	Author     *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// middleware
func (c *Course) IsEmpty() bool {
	return c.CourseID == "" || c.CourseName == ""
}

func main() {
	fmt.Println("API in Go")
}
