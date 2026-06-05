// Ok let's do this! Chapter 8 is Arrays. Let's talk about them.

package main
// We need it! Without it the complier will complain ...expected 'package', found 'import' ...

import "fmt"
// Remember, fmt stands for the Format package. It contains function to print to the console. We need it too.

func main() {
	// We need func main() to. Without it the complier will complain ...syntax error: non-declaration statement outside function body

	// Finally we can start. Arrays are fixed length sequences of elements of the same type.
	// Arrays are not the same as arrays in other languages.
	// Length is a part of their type and hence they cannot be resized. Weird, I know.
	// Here is a basic example.

	var a = [5]int{1, 2, 3, 4, 5}
	fmt.Println("a: ", a)
}