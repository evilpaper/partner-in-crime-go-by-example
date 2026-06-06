// Ok let's do this! Chapter 9 is Slices. Let's talk about them.

package main

import "fmt"

func main() {

	// Ok, lets do this. We talk about slices.

	// Slices are like arrays but more flexible.

	// Unlike arrays, slices are typed only by the elements they contain (not the number of elements). 

	// Here is an uninitialized slice. It equals to nil and has length 0.
	var s []string
	fmt.Println("uninit slice:", s, s == nil, len(s) == 0)

	// Create a slice with make.
	// What is make? It is a built-in function that creates a new slice.
	// Can I use make for other things? Yes, you can use make for other things like maps and channels.
	// The make() function is one of Go's built-in tools used primarily for memory allocation and initialization.
	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	// Create a slice with some values.
	var pizzaSlices = []string{"Pepperoni", "Road to Napoli", "Goodfellas"}

	// We can set and get just like with arrays.
	pizzaSlices[2] = "Donnie Brasco"
	
	fmt.Println("pizzaSlices:", pizzaSlices)
}