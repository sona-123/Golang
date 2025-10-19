package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "sonali"
	fmt.Println(s)
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println(s)
	fmt.Println("Rune count:", utf8.RuneCountInString(s))
	//In Go, a string is a read-only slice of bytes.
	/*Each character in a Go string is a byte (an alias for uint8).
	Strings are immutable — you can’t modify them once created.
	Internally, they store UTF-8 encoded bytes, not necessarily one byte per character. */
	s1 := "हाय"
	fmt.Println(s1) // 9 (since each Devanagari character = 3 bytes)

}
