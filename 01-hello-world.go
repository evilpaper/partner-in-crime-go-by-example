// Our first program will print the classic "hello world"
// message. Here's the full source code.

// All Go files start with a package declaration.
// main package is the entry point of the program.
package main

// import packages we need.
import "fmt"

// main function is the entry point of the program.
// It's the first function that is executed when the program runs.
func main() {
	// Println is a function that prints a string to the console.
	fmt.Println("Hello world!")
}