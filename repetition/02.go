package main

import "fmt"

func main() {

	// Ok the main setup is done. What should we do here? We talk about values right?
	// Yes! Go has values. The regular values like strings, integers, floats, booleans, and more.
	// Let's print some of them and see what they are.

	// Strings. Can be added together with the + operator.
	fmt.Println("go" + "lang")
	// Integers. Whole numbers.
	fmt.Println("1+1 =", 1+1)
	// Floats. Decimal numbers.
	fmt.Println("7.0/3.0 =", 7.0/3.0)
	// Booleans. True or false. With normal boolean operators you expect.
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}