package main

import "fmt"

func main() {

	// Ok, lets do this. We talk about switch statements.

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
}