package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	res := add(1, 2)
	fmt.Println(res)

	fmt.Println("result is: ", sub(4, 1))
	fmt.Println(proAdder(3, 4, 5, 6, 7, 9))
}

func add(x int, y int) int {
	return x + y
}

func sub(x int, y int) int {
	return x - y
}

func proAdder(values ...int) int {
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}
