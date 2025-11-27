package main

import "fmt"

// An enum is a type that has a fixed number of possible values, each with a distinct name.
// Go doesn’t have an enum type as a distinct language feature.
// But enums are simple to implement using existing language idioms.

// Our enum type ServerState has an underlying int type.
type ServerState int

const (
    StateIdle ServerState = iota // StateIdle now has a value of 0. Wtf! 
    StateConnected
    StateError
    StateRetrying
)

var stateName = map[ServerState]string{
    StateIdle:      "idle",
    StateConnected: "connected",
    StateError:     "error",
    StateRetrying:  "retrying",
}

// String() method implements fmt.Stringer interface, overriding default string representation
// This makes fmt.Println(StateIdle) print "idle" instead of "0"
// Funky name, Stringer. You can read more about it here: https://go.dev/tour/methods/17

func (ss ServerState) String() string {
    return stateName[ss]
}

func main() {
    fmt.Println(StateIdle) // Prints idle and not 0 since we overrode the default string representation. Wtf again!

    ns := transition(StateIdle)
    fmt.Println(ns)

    ns2 := transition(ns)
    fmt.Println(ns2)
}

func transition(s ServerState) ServerState {
    switch s {
    case StateIdle:
        return StateConnected
    case StateConnected, StateRetrying:

        return StateIdle
    case StateError:
        return StateError
    default:
        panic(fmt.Errorf("unknown state: %s", s))
    }
}

/**

wtf!

Go has a lot of “clever laziness”

When you write a const block like this:

const (
    A = 1
    B
    C
)

Go automatically fills in the blanks for you.
It sees that B and C don’t have values, so it says: “okay, B and C will be 1 and 1.”
The key rule: “If you don’t specify a value, reuse the previous expression.”

iota changes with each new line.

So we have:

const (
    A = iota  // iota = 0
    B         // iota = 1
    C         // iota = 2
)

In short

Go doesn’t magically know to assign numbers — it just:
Repeats the previous expression if you don’t write a new one
iota is Go’s auto-counter for constants.
Increments iota each line in the block

Boom — automatic, elegant counting!

iota is a Greek letter (ι) that basically means “the smallest thing” or “a tiny amount.”
You can remember it by thinking “iota” > “I Ought’a Count” 

wtf again!

- When fmt.Println() tries to print StateIdle, it looks for a String() 
method on the ServerState type

- Go finds the String() method defined on lines 25-27

- This String() method overides the default behavior and maps StateIdle to "idle" instead of "0"

*/