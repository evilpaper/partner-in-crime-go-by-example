package main

import "fmt"

// 1. Structs are collections of fields.

// 2. We can define a struct by using the type keyword and then the struct name and then the fields.

type person struct {
	name string
	age int
}

func main() {
	fmt.Println(person{"John", 30})
}