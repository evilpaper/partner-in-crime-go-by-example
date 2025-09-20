package main

import "fmt"

// Go support pointers, allowing you to pass references to values and records within your program.
// A pointer holds the memory address of a value.
// The type *T is a pointer to a T value. Its zero value is nil.
// The * operator denotes the pointer's underlying value.
// The & operator generates a pointer to its operand. 

// Pointers and Values are different things.
// zeroval has an int parameter, so arguments will be passed to it by value.
// zeroval will get a copy of ival distinct from the one in the calling function.
func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	// i is 1, just like that
	i := 1
	fmt.Println("initial:", i)

	// zeroval gets a copy of ival distinct from the one in the calling function.
	// so i is still 1. 
	// so zeroval doesn't change the i in the calling function.
	// Don't get confused by the copy of ival and i.
	zeroval(i)
	fmt.Println("zeroval:", i)

	// zeroptr has an *int parameter, so arguments will be passed to it by reference.
	// zeroptr will get a pointer to the i in the calling function.
	// so zeroptr can change the i in the calling function.
	// Don't get confused by the pointer of iptr and i.
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}