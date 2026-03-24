package main

import "fmt"

// Methods are functions that are associated with a type.
// Just like methods in JavaScript.

// 1. First we need a struct. Like an object in JavaScript.

type rect struct {
	width, height int
}

// 2. Then we can define a method on the struct. Just like a method in JavaScript.

func (r *rect) area() int {
	return r.width * r.height
}

func main() {

	r := rect{width: 10, height: 5}
	fmt.Println(r.area())

}