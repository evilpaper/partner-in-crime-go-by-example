package main
import (
	"fmt"
)

func main() {

	// Slices are an important data type in Go, giving a more powerful interface to sequences than arrays.
	
    // Slices are types by the elements they contain.
	// Funky, similar syntax to arrays.
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	// Create slice with non-zero length with make.
	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))
	// len = length, cap = capacity

	// Set and get values.
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// len returns the length
	fmt.Println("len:", len(s))

	// append returns a new slice with the new element added.
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// copy copies elements from one slice to another.
	c := make([]string,len(s))
	copy(c, s)
	fmt.Println("cpy:", c)


}