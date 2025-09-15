package main

import "fmt"

// Functions are central in Go.

// Take two ints and return an int.
func plus(a int, b int) int {
	return a + b
}

// Go requires explicit returns, i.e. it won't automatically return the value of the last expression.
func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	// Call as expected, with name(args).
	res := plus(1, 2)
	fmt.Println(" 1 + 2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println(" 1 + 2+ 3 =", res)
}