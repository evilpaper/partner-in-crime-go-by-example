package main
import (
	"fmt"
)

func main() {

	// Slices are an important data type in Go, giving a more powerful interface to sequences than arrays.
	
    // Slices are types by the elements they contain.
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)
}