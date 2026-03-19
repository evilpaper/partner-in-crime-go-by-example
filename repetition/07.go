package main

import (
	"fmt"
	"time"
)

func main() {

	// Ok, lets do this. We talk about switch statements.

	// A basic switch statement
	var i = 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// We can use comma to separate multiple cases.
	var weekday = time.Now().Weekday()
	switch weekday {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}
}