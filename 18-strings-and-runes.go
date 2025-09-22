package main

import (
	"fmt"
	"unicode/utf8"
)

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

	// Indexing into a string produces the raw byte values 
	// at each index. This loop generates the hex values of
	// all the bytes that constitute the code points in s.
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	// To count how many runes are in a string, we can use the utf8 package.
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	// Range loop decodes each UTF-8 sequence and provides the start position of the rune.
	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}
}