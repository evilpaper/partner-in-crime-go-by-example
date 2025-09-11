// Arrays are fixed length sequences of elements of the same type.
// Slices are more common.

package main

import "fmt"

func main() {

	// An array that hold 5 ints.
	var a [5]int
	fmt.Println("emp:", a)

	// Set the value at index 4 to 100.
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// Get the length of the array.
	fmt.Println("len:", len(a))

	// Declare and initialize an array.
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)
}