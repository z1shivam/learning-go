package main

import "fmt"

func main() {
	fmt.Println("welcome to array")

	var fruitList [4]string
	fruitList[0] = "Mango"
	fruitList[1] = "Apple"
	fruitList[2] = "Banana"

	fmt.Println("fruitList is: ", fruitList)
	fmt.Println("length of the array: ", len(fruitList))

	var vegList = [4]string{"potato", "mushroom", "spinach", "something"}
	// prints 5 even though there are only 3 elements because it prints the capacity and not the number of current element.

	fmt.Println(vegList, len(vegList))
}
