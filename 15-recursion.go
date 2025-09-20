package main

import "fmt"

// Go supports recursive functions.

// This function calls itself until it reaches the base case.
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))
}