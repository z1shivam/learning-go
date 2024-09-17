package main

import "fmt"

// defining a closure. it is basically a function returning a function. kind of anonymous function.
func callThisToGetFunction() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	thisFuncWillBeRef := callThisToGetFunction()

	fmt.Println(thisFuncWillBeRef())
	fmt.Println(thisFuncWillBeRef())
	fmt.Println(thisFuncWillBeRef())
}
