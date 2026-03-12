// Our first program will print the classic "hello world" message.

// All Go files start with a package declaration.
// The main package marks the entry point of the program.
package main

// We then import packages we need.
import "fmt"

// We can then declare the main function. The entry point of the program.
// It's the function the Go compiler will look for first when the program runs.
func main() {
	// Println is a function that prints a string to the console.
	fmt.Println("Hello world!")
}

// ----------------------------------
// This is a program that prints "Hello world!" to the console.
// package main — Declares this is an executable program (not a library). It tells Go this package can be built into a runnable binary.
// func main() — The entry point where execution starts. When you run the program, Go calls main().
// Analogy: package main is like labeling a box "This is a complete toy" (not just a part). func main() is like the "ON" button — pressing it starts the toy.
// In practice:
// Without package main, Go treats it as a library (can't run it directly).
// Without func main(), Go doesn't know where to start (compilation error).