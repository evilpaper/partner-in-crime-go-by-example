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

	// Let the compiler count the number of elements.
	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

    // [...]int means: "Create an array of type int, and let Go figure out the length automatically based on the initializer."
	// 3: 400 → this explicitly sets the element at index 3 to 400.
	// the index before 3 is zeroed (will be 0)
	b = [...]int{100, 3: 400, 500}
	fmt.Println("idx:", b)

	// Array are one dimensional, but can be composed to form multidimensional arrays.
	// First an array with two items, each item is an array with three items.
	var twoD [2][3]int
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	var bulletTrain [2][3]int
	fmt.Println("bulletTrain: ", bulletTrain)

   // Initialize a 2D array with a shorter syntax.
   twoD = [2][3]int{
	{1, 2, 3},
	{1, 2, 3},
   }
   fmt.Println("2d: ", twoD)
}