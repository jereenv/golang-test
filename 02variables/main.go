package main

import "fmt"


// no walrus assignment outside function

// jwt := 3000

const LoginToken string = "sdasfasdasd"  //is plublic

func main() {
	var username string = "Jereen"
	fmt.Println(username)
	fmt.Printf("Variable is of type; %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type; %T \n", isLoggedIn)

	var smallVal uint8 = 255 
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type; %T \n", smallVal)


	var smallFloat float64 = 255.12312313213143 
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type; %T \n", smallFloat)


	// default values and some aliases

	var anotherVariable string
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type; %T \n", anotherVariable)


	// implicit type

	var website = "github.com"
	fmt.Println(website)


	// no var style

	numberofUser := 3000
	fmt.Println(numberofUser)


	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type; %T \n", LoginToken)




}
