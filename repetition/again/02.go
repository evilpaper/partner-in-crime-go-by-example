// Ok let's do this! Chapter 2 is values. Let's talk about them.

package main
// We need it! Without it the complier will complain ...expected 'package', found 'import'

import "fmt"
// Remember, fmt stands for the Format package. It contains function to print to the console. We need it too.

func main() {
	// We need func main() to. Without it the complier will complain ...syntax error: non-declaration statement outside function body
	fmt.Println("I am " + "two strings" + " added together")
	// Finally, now we can print. Go has the usual suspect of values.
	// Here are strings concatenated together.
	fmt.Println("1+1 =", 1+1)
	// Here are numbers added together.
	fmt.Println("7.0/3.0 =", 7.0/3.0)
	// Here are floats divided together.
	fmt.Println(true && false)
	// Here are booleans anded together.
	fmt.Println(true || false)
	// Here are booleans ored together.
}

// Done! As ususal run it from the command line with go run <path to the file>
