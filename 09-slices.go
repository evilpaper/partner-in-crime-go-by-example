package main

import (
	"fmt"
	"slices"
)

// Slices are an important data type in Go, giving a more powerful interface to sequences than arrays.

func main() {

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

	// slice syntax
	l := s[2:5]
	fmt.Println("sl1:", l)

	// slice up to (but excluding) index 5
	l = s[:5]

	// slice up from (and including) index 2
	l = s[2:]
	fmt.Println("sl3:", l)

	// declare and initialize a slice in one line
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// slice package contains functions to manipulate slices
	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t and t2 are equal")
	}

	// multi-dimensional slices
	twoD := make ([][]int, 3)
	for i := range 3 {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := range innerLen {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}