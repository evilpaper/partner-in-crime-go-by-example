package main

import "fmt"

func returnTwoValues() (int, int) {
	return 1, 2
}

func main() {

	// Ok, lets do this. We talk about multiple return values.
	// What are they?
	// Just the fact that a function can return multiple values.
	// Like in JavaScript.
	// Let's see an example.

	// We define a function that returns two values.
	a, b := returnTwoValues()
	fmt.Println(a, b)
	// We call the function and assign the returned values to two variables.
	
}