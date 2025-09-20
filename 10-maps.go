package main

import (
	"fmt"
	"maps"
)

// Maps are key-value pairs.
// Maps are Go's built-in associative data type
// Called hashes or dicts in other languages.

func main() {

	// Create a new map with make.
	m := make(map[string]int)

	// Set key/value pairs using typical name[key] = val syntax.
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	// Get a value with name[key].
	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	// Return zero value if key doesn't exist.
	fmt.Println("v3: ", m["k3"])

	// len returns the number of key/value pairs.
	fmt.Println("len:", len(m))

	// Delete a key/value pair with delete.
	delete(m, "k2")
	fmt.Println("map:", m)

	// Check if a key exists with a two-value assignment.
	// The optional second return value when getting a value from a map indicates if the key was present in the map.
	_, prs := m["k2"]
    fmt.Println("prs:", prs)

	// Declare and initialize a map in a single line.
	n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n)

	// maps package contains a number of useful utility functions.
	n2 := map[string]int{"foo": 1, "bar": 2}
    if maps.Equal(n, n2) {
        fmt.Println("n == n2")
    }
}