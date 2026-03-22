package main

import "fmt"

// 1. Structs are collections of fields.

// 2. We can define a struct by using...
// 2.1 the type keyword then, 
// 2.2 the name of the struct then,
// 2.3 the struct keyword and then,
// 2.4 the fields.

type person struct {
	name string
	age int
}

func main() {
	fmt.Println(person{"John", 30})
}