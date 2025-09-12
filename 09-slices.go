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
	

}