package main

import "fmt"

// for is Go’s only looping construct

func main() {
 
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

	// Another way, “do this N times”
	for i := range 3 {
		fmt.Println("range", i)
	}

	// Without a condition, it loops forever until break or return
	for {
		fmt.Println("loop")
		break
	}

	// continue to the next iteration of the loop
	for n := range 6 {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
    }

}