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

	// Range provide and both the index and value.
	for i, num := range nums {
		if num  == 3 {
			fmt.Println("index:", i)
		}
	}

	
	kvs := map[string]string{"a": "apple", "b": "banana"}
	// Range on map iterates over key/value pairs.
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// Range over just keys
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// Range over strings iterates over unicode code points.
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}