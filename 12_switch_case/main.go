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
	switch randomNum {
	case 1:
		fmt.Println("dice value is 1")
		fallthrough // it will fall through to next case, but there is no fallthrough on next case, so it will stop there.
	case 2:
		fmt.Println("dice value is 2")
	case 3:
		fmt.Println("dice value is 3")
	default:
		fmt.Println("what was that")
	}
}
