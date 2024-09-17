package main

import "fmt"

func main() {
	fmt.Println("welcome to loops...")

	days := []string{"sunday", "tuesday", "wednesday", "friday", "saturday"}

	fmt.Println(days)

	// type 1 for loop
	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	// type 2 for loop
	for i := range days {
		fmt.Println("type 2 for loop", days[i])
	}

	// type 3 for loop
	for _, day := range days {
		fmt.Println("type 3 for loop", day)
	}

	rogueValue := 1
	for rogueValue < 10 {

		fmt.Println("rogueValue is", rogueValue)
		rogueValue++
		if rogueValue == 5 {
			break // there is also continue (same concept)
		}
		goto shivam
	}

	// creating a label
shivam:
	fmt.Println("this is a jump statement")
}
