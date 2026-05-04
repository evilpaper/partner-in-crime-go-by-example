package main

import "fmt"

func fact(n int) int {
	// Start with the base case.
	if n == 0 {
		return 1
	}
	// All other cases we call the function again.
	return n * fact(n-1)
}

func main() {

	// Here we go. We talk about recursion.

	// Let's say we want to calculate the factorial of a number.
	// Definition of factorial: The factorial function multiplies a number by every whole number below it down to 1.
	// Example: 5!=5×4×3×2×1
	fmt.Println(fact(7))
}