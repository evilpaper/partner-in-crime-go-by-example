package main

import "fmt"

func main() {
	var a = "initial" // use var to declare a variable
	fmt.Println(a)

	var b, c int = 1, 2 // declare multiple variables
	fmt.Println(b, c)

	var d = true // go will infer the type of the variable
	fmt.Println(d)
	
	var e int // declare a variable without initializing it, gets the zero value
	fmt.Println(e)

	f := "apple" // short assignment statement, declares and initializes a variable
	fmt.Println(f)

	var g = "pear" // same thing as :=
	fmt.Println(g)
}