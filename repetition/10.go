package main

import "fmt"

func main() {

	// Ok, lets do this. We talk about maps.

	// Maps in Go are actually very similar to the Map in JavaScript.

	// Creating a map
	userScores := map[string]int{
		"Alice": 95,
		"Bob": 82,
		"Charlie": 90,
	}

	fmt.Println(userScores)
}