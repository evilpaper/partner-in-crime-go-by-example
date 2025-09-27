package main

import "fmt"

// Go supports methods defined on struct types.

type rect struct {
	width, height int
}

// This are method has a receiver type of *rect (a pointer to rect).
// Remember * is the pointer operator.
func (r *rect) area() int {
	return r.width * r.height
}

// This method has a receiver that is a value of type rect.
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("Area:", r.area())
	fmt.Println("Perimeter:", r.perim())

	// Create a pointer to the rect
	// Go automatically handles conversion 
	// between values and pointers for method calls. 
	// Why do it? => You may want to use a pointer receiver type 
	// to avoid copying on method calls or to allow 
	// the method to mutate the receiving struct
	rp := &r
	fmt.Println("Area:", rp.area())
	fmt.Println("Perimeter:", rp.perim())
}