package main

import "fmt"

func main() {

    fmt.Println("go" + "lang")

    fmt.Println("1+1 =", 1+1)
    fmt.Println("7.0/3.0 =", 7.0/3.0)

    fmt.Println(true && false)
    fmt.Println(true || false)
    fmt.Println(!true)
}

// Explain this like I am 5 years old
// ----------------------------------
// This program is like a magic box that shows us different types of things:
//
// 1. String addition (line 7): We're putting two words together!
//    "go" + "lang" = "golang" (like putting two toy blocks together)
//
// 2. Math with numbers (lines 9-10): We're doing simple math!
//    1 + 1 = 2 (just like counting on your fingers)
//    7.0 / 3.0 = 2.333... (dividing is like sharing cookies - you get pieces!)
//
// 3. True/False questions (lines 12-14): We're asking yes/no questions!
//    true && false = false (both need to be true, like needing BOTH hands to clap)
//    true || false = true (only one needs to be true, like needing ONE hand to wave)
//    !true = false (the ! means "not" - so "not true" means false, like turning a light switch off)
