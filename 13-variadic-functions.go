package main

// Variadic functions can be called with any number of trailing arguments.
// For example, fmt.Println is a common variadic function.

import "fmt"

// A function that take an arbitrary number of ints as arguments.
func sum(nums ...int) {
	fmt.Println("------ Let's go! ------")
	fmt.Println(nums, " ")

	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println("total: ", total)
	fmt.Println("number of args: ", len(nums))
}

func main() {

	// Variadic functions can be called in the usual way with individual arguments.
	sum(1, 2)
	sum(1, 2, 3)
	sum(1, 2, 3, 4)

	// If you already have multiple args in a slice, apply them to a variadic function using func(slice...) .
	nums := []int{1, 2, 3, 4, 5}
	sum(nums...)
}

