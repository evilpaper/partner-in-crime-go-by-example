package main

import "fmt"

func main() {

	// Ok, lets do this. We talk about arrays. 

	// Arrays are fixed length sequences of elements of the same type.

	// Array in Go are not the same as arrays in other languages.

	// Length is a part of their type and hence they cannot be resized. Weird, I know.

	// Here is a basic example.
	var a = [5]int{1, 2, 3, 4, 5}
	fmt.Println("a: ", a)

	// There isn’t a push method because arrays are of fixed lengths in Go
	
}