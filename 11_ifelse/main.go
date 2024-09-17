package main

import "fmt"

func main() {
	fmt.Println("if else in go")

	loginCount := 23
	var result string

	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount > 10 {
		result = "New User"
	} else {
		result = "watchout"
	}

	if true {
		fmt.Println("this will always be printed...")
	}

	// you can declare variable on the go
	// and remember that there is no ternary in go. so you will have to use if else even for basic tasks.

	if num := 3; num < 10 {
		fmt.Println("number is less than 10")
	}

	fmt.Println(result)
}
