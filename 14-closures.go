package main

// Go supports anonymous functions, which can form closures.
// Anonymous functions are useful when you want to define a function inline without giving it a name.

import "fmt"

// This function returns another function.
// The returned function closes over the variable i to form a closure.
func inSeq() func() int{
	 i:= 0
	 return func() int {
		i++
		return i
	 }
}

func main() {
	// We call inSeq, assigning the result (a function) to nextInt.
	// This function value captures its own i value, which will be updated each time we call nextInt.
	nextInt := inSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
}