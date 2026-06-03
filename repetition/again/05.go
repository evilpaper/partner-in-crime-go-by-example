// Ok let's do this! Chapter 5 is loops. Let's talk about them.

package main
// We need it! Without it the complier will complain ...expected 'package', found 'import' ...

import "fmt"
// Remember, fmt stands for the Format package. It contains function to print to the console. We need it too.

func main() {
	// For is Go's only looping construct
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
}