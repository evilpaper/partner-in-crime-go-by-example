package main

import "fmt"

func main() {

	// Ok, lets do this. We talk about loops.

	// For is Go's only looping construct

	// Basic loop, single condition
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// Classic init/condition/after syntax
	for j := 0; j <=3; j++ {
		fmt.Println(j)
	}
}