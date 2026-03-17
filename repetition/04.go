package main

import "fmt"

func main() {

	// Ok, lets do this. We talk about constants.
	
	// Go support constants. Constants are like variables, but they cannot be changed.
	
	// Constants are declared using the const keyword.
	const s string = "I am a constant"

	// Constants can be string, boolean, and numeric values.
	const n = 500000000

	fmt.Println(s)
	fmt.Println(n)
}