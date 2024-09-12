package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time module!!")

	presentTime := time.Now().Format("01-02-2006 15:04:05 Mon")

	fmt.Println(presentTime)

	createdDate := time.Date(2024, time.October, 4, 12, 12, 12, 12, time.Local)
	fmt.Println(createdDate.Format("02-01-06"))
}
