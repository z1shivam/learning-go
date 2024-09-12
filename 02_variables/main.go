package main

import "fmt"

func main() {
	var username string = "Shivam"
	fmt.Println("username", username)
	fmt.Printf("variable is of type %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("%T", isLoggedIn)
}
