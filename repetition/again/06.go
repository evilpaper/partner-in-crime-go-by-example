// Ok let's do this! Chapter 6 is If/Else. Let's talk about it.

package main
// We need it! Without it the complier will complain ...expected 'package', found 'import' ...

import "fmt"
// Remember, fmt stands for the Format package. It contains function to print to the console. We need it too.

func main() {
	// We need func main() to. Without it the complier will complain ...syntax error: non-declaration statement outside function body

	// Finally we can start. If/Else is used for branching.
	// Here is a basic example.
	
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// Logical operators like && and || are often useful in conditions.
	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("Either 8 or 7 are even. Guess which?")
	}
}

// Done! As ususal run it from the command line with go run <path to the file>