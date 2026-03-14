package main

import "fmt"

func main() {

	// Ok the main setup is done. What should we do here? We talk about variables right?

	// First thing first. We declare variables with the var keyword.
	var a = "initial"
	fmt.Println(a)

	// We can declare multiple variables at once.
	var b, c int = 1, 2
	fmt.Println(b, c)

	// We can declare a variable without initializing it.
	var d int
	fmt.Println(d)

	// We can declare a variable with a short assignment statement.
	f := "apple"
	fmt.Println(f)
}