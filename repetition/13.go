package main

import "fmt"

func sum(nums ...int) {
	fmt.Println(nums, " ")

	total := 0
	for _, num := range nums {
		total += num
	}
}
func main() {

	// Ok, lets do this. We talk about variadic functions.

	// Variadic functions are functions that can take a variable number of arguments.
	// They are declared with the ... type suffix.

	// Here is a basic example.

	sum(1, 2, 3)
	// sum(1, 2, 3, 4)
	// sum(1, 2, 3, 4, 5)

	// Here is a more complex example.

	sum(1, 2, 3, 4, 5)

	// Here is a more complex example.

	sum(1, 2, 3, 4, 5)
}