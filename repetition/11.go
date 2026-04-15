package main

import "fmt"

func greet(name string) {
	fmt.Println("Hello, " + name)
}

func main() {

	// Ok, lets do this. We talk about functions.

	// Functions are pretty similar to how they work in JavaScript.
	// Their signature is also pretty similar to JavaScript apart from the func keyword instead of function.
	// Seems we have to declare them outside of the main function.
	greet("John")
	greet("Jane")
	greet("Jim")
	greet("Jill")
	greet("Jack")
	greet("Jill")
}

