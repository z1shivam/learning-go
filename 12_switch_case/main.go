package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("switch case in golang")

	randomNum := rand.Intn(6) + 1
	fmt.Println(randomNum)

	if randomNum == 1 {
		fmt.Println("Yay, you can start the game!!")
	}

	// there is no automatic fall through, so there is no need for break; to make it do so, you can use fallthrough
	switch randomNum { // switch can be exist without a variable. it will be alternative to ifelse
	case 1:
		fmt.Println("dice value is 1")
		fallthrough // it will fall through to next case, but there is no fallthrough on next case, so it will stop there.
	case 2:
		fmt.Println("dice value is 2")
	case 3, 4, 5: // you can write multiple cases separated by commas.
		fmt.Println("dice value is 3")
	default:
		fmt.Println("what was that")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}

// A type switch compares types instead of values. You can use this to discover the type of an interface value. In this example, the variable t will have the type corresponding to its clause.
