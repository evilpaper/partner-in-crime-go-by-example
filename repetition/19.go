package main

import "fmt"

// Structs are collections of fields. Other languages call them classes or objects. JavaScript calls them objects.

type person struct {
	name string
	age int
}

func main() {
	fmt.Println(person{"Bob", 20})
}