package main

import "fmt"

func main() {
	var username string = "Shivam"
	fmt.Println("username", username)
	fmt.Printf("variable is of type %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("%T", isLoggedIn)

	var smallFloat float32 = 473.49839 // only 5 value after the decimal as precesion
	fmt.Println(smallFloat)

	// data types aliases & default values
	var anotherVar int
	fmt.Println(anotherVar)

	var website = "sivam"
	fmt.Println(website)

	// no var style
	numberOfUsers := 32
	fmt.Println(numberOfUsers)
}

// The first letter of the variable makes it public and accessible to all of program.
