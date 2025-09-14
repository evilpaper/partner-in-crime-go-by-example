package main

import "fmt"

func main() {
	// Maps are key-value pairs.
	// Maps are Go's built-in associative data type
	// Called hashes or dicts in other languages.

	// Create a new map with make.
	m := make(map[string]int)

	// Set key/value pairs using typical name[key] = val syntax.
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	// Get a value with name[key].
	v1 := m["k1"]
	fmt.Println("v1: ", v1)
}