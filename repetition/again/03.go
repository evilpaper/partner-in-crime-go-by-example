// Ok let's do this! Chapter 3 is variables. Let's talk about them.

package main
// We need it! Without it the complier will complain ...expected 'package', found 'import' ...

import "fmt"
// Remember, fmt stands for the Format package. It contains function to print to the console. We need it too.

func main() {
	// We need func main() to. Without it the complier will complain ...syntax error: non-declaration statement outside function body

	// Finally we can start. Variables hold values of a certain type.
	// We declare variables with the var keyword.
	var a = "initial"
	fmt.Println(a)

	// We can declare multiple variables at once.
	var b, c int = 1, 2
	fmt.Println(b, c)
}
