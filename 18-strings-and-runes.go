package main

import "fmt"

// A Go string is a read-only slice of bytes.
// Strings are treated as containers of text encoded in UTF-8.
// In other languages, strings are a sequence of characters.
// In Go, strings are a sequence of bytes.
// In Go, the concept of a character is called a rune.
// It’s an integer that represents a Unicode code point.
// Go strings are immutable.

func main() {
	// s is a string assigned a literal value representing 
	// the word “hello” in the Thai language.
	 const s = "สวัสดี"

	// Since strings are equivalent to []byte, this will
	// produce the length of the raw bytes stored within.
	fmt.Println("len:", len(s))
}