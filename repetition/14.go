package main

import "fmt"

func inSequence() func() int {
	i := 0
	// This is the closure.
	return func() int {
		i++
		return i // This is the return value of the closure.
	}
}

func main() {

	// Ok, lets do this. We talk closures.

	// Let's say we want nextInt to contain that return numbers in sequence.
	// 1, 2, 3, 4, 5, ... you get the idea.
	// We set it up like this:
	nextInt := inSequence()
	// Then we call it like this:
	fmt.Println(nextInt()) // Since inSequence returns a function, we can call it like this.
	fmt.Println(nextInt()) // And we get the next number in the sequence.
	fmt.Println(nextInt()) // And we get the next number in the sequence.
}