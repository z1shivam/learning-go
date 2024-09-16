package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("heyy")

	var mySlice = []string{}

	mySlice = append(mySlice, "firstItem", "secondItem", "thirdItem")
	mySlice = append(mySlice[2:], "fourthItem")
	mySlice = mySlice[1:]

	highScores := make([]int, 4)
	highScores[0] = 23
	highScores[1] = 92
	highScores[2] = 34
	highScores[3] = 43
	// highScores[4] = 55

	sort.Ints(highScores)

	var courses = []string{"reactjs", "python", "ruby", "cplusplus", "javascript"}
	fmt.Println(courses)

	courses = append(courses[:3], courses[4:]...)

	fmt.Println(courses)

	fmt.Println(highScores)
	fmt.Println(mySlice)
}
