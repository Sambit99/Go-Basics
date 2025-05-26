package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const getUrl = "https://jsonplaceholder.typicode.com/todos/1"
const postUrl = "https://jsonplaceholder.typicode.com/posts"

func main() {
	fmt.Println("Get Request")

	PerformGetRequest(getUrl)
	PerformPostJsonRequest(postUrl)
	PerformPostFormRequest()
}

func PerformGetRequest(url string) {

	response,err := http.Get(url)

	if err != nil{
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code: ",response.StatusCode)
	fmt.Println("content length is: ",response.ContentLength)

	var responseString strings.Builder
	content,_ := io.ReadAll(response.Body)
	byteCount,_ := responseString.Write(content)

	// fmt.Println("Content: ",string(content))
	fmt.Println(byteCount)
	fmt.Println(responseString.String())
}

func PerformPostJsonRequest(url string) {
	requestBody := strings.NewReader(`
		{
			"userId": 1,
			"id": 1,
			"title": "Dummy Title",
			"body": "Dummy Data"
		}
	`)

	response,err := http.Post(url,"application/json",requestBody)

	if err != nil{
		panic(err)
	}

	defer response.Body.Close()

	content,_ := io.ReadAll(response.Body)

	fmt.Println("Content: ",string(content))
}

func PerformPostFormRequest(){
	const myUrl = "http://localhost:3000/api/v1/auth/signup"

	data := url.Values{}

	data.Add("username","Sambit26")
	data.Add("email","sam26@gmail.com")
	data.Add("fullname","Sambit26")
    data.Add("password" , "123456789")

	response,err := http.PostForm(myUrl, data)

	if err != nil{
		panic(err)
	}

	defer response.Body.Close()

	content,_ := io.ReadAll(response.Body)

	fmt.Println("Content: ",string(content))

}