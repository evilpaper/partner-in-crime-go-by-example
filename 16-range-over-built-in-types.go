package main

import "fmt"

func main() {
	// Use range to sum the numbers in a slice.
	nums := []int{2,3,4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
}