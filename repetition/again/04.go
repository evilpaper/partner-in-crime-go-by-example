// Ok let's do this! Chapter 4 is constants. Let's talk about them.

package main
// We need it! Without it the complier will complain ...expected 'package', found 'import' ...

import "fmt"
// Remember, fmt stands for the Format package. It contains function to print to the console. We need it too.

func main() {
	// Constants are declared using the const keyword.
	const s string = "I am a constant"

	// Constants can be string, boolean, and numeric values.
	const n = 500000000
	
	fmt.Println(s)
	fmt.Println(n)
}