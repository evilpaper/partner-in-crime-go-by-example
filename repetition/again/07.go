// Ok let's do this! Chapter 7 is Switch. Let's talk about it.

package main
// We need it! Without it the complier will complain ...expected 'package', found 'import' ...

import (
	"fmt"
	"time"
)
// Remember, fmt stands for the Format package. It contains function to print to the console. We need it too.

func main() {
	// We need func main() to. Without it the complier will complain ...syntax error: non-declaration statement outside function body

	// Finally we can start. Switch is used for branching.
	// A basic switch statement
	var i = 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// We can use comma to separate multiple cases.
	var weekday = time.Now().Weekday()
	fmt.Println("Today is", weekday)
	switch weekday {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}
}