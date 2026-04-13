package main

import "fmt"

func main() {

	// Ok, lets do this. We talk about slices.

	// Slices are like arrays but more flexible.

	// Unlike arrays, slices are typed only by the elements they contain (not the number of elements). 
	// An uninitialized slice equals to nil and has length 0.
	var s []string
	fmt.Println("uninit slice:", s, s == nil, len(s) == 0)
}