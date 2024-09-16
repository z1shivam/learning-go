package main

import "fmt"

func main() {
	fmt.Println("structs in golang...")

	// no inheritance in golang; no super no parent.
	var shivam = User{"Shivam Kumar", "shivamk@gmail.com", true, 21}
	fmt.Println(shivam)

	fmt.Printf("shviam details are: %+v\n", shivam)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
