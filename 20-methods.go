package main

import "fmt"

// Go supports methods defined on struct types.

type rect struct {
	width, height int
}

// This area method has a receiver type of rect.
func (r rect) area() int {
	return r.width * r.height
}

// This method has a receiver type of *rect.
func (r *rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}
	
	fmt.Println("Area:", r.area())
	fmt.Println("Perimeter:", r.perim())
}