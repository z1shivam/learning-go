package main

import "fmt"

func main() {
	var shivam *int
	fmt.Println(shivam)

	dataValue := 23
	var myPtr = &dataValue

	fmt.Println(*myPtr)
}
