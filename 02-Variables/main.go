package main

import "fmt"

const JwtToken string = "qcwh"; // Public

func main() {
	var username string = "sambit";
	fmt.Println(username);
	fmt.Printf("Variable is of type : %T \n",username);


	var isLoggedIn bool = false;
	fmt.Println(isLoggedIn);
	fmt.Printf("Variable is of type : %T \n",isLoggedIn);

	var smallVal uint8 = 255;
	fmt.Println(smallVal);
	fmt.Printf("Variable is of type : %T \n",smallVal);

	var smallFloat float32 = 255.3453451234141;
	fmt.Println(smallFloat);
	fmt.Printf("Variable is of type : %T \n",smallFloat);

	var longFloat float64 = 255.3453451234141;
	fmt.Println(longFloat);
	fmt.Printf("Variable is of type : %T \n",longFloat);

	// Derfault Values and aliases

	var anotherVariable int;
	fmt.Println(anotherVariable);
	fmt.Printf("Variable is of type : %T \n",anotherVariable);

	//  Implicit Type

	var website = "Hello World";
	fmt.Println(website);
	fmt.Printf("Variable is of type : %T \n",website);

	// No Var Style

	numberOfUsers := 30000;
	fmt.Println(numberOfUsers);
	fmt.Printf("Variable is of type : %T \n",numberOfUsers);

	fmt.Println(JwtToken);
	fmt.Printf("Variable is of type : %T \n",JwtToken);

}