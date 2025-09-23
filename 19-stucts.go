package main

import "fmt"

// Go’s structs are typed collections of fields.

// This person struct type has name and age fields.
type person struct {
	name string
	age int
}

func main() {

	// This syntax creates a new struct.
	fmt.Println(person{"Bob", 20})

	// Name the fields when initializing a struct.
	fmt.Println(person{name: "Alice", age: 30})

	// Omitted fields will be zero-valued.
	fmt.Println(person{name: "Fred"})

}