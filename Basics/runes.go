package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	/*A rune in Go is an alias for int32, representing a Unicode code point (one character).
	  To correctly handle characters, convert the string to a slice of runes:*/

	//String to a slice of runes
	runes := []rune("sonali")
	fmt.Println(runes)
	fmt.Println(len(runes))

	for i, r := range runes {
		fmt.Println("%d : %c\n", i, r)
	}

	//Range loop with strings
	s := "sonali"
	for i, r := range s {
		/*i is the byte index of the rune in the string.
		r is the rune value (Unicode code point).*/
		fmt.Println("index: %d, rune: %c\n", i, r)
	}

	/*String ↔ Rune/Byte Conversions
	Conversion	Example	Meaning
	[]byte(s)	b := []byte("Go")	Get raw UTF-8 bytes
	string(b)	s := string([]byte{71, 111})	Convert bytes to string
	[]rune(s)	r := []rune("हाय")	Decode into Unicode runes
	string(r)	s := string([]rune{2361, 2366, 2351})	Encode runes into string*/

	//String immutability

	s1 := "hello"
	r := []rune(s1)
	r[0] = 'H'
	s1 = string(r)
	fmt.Println(s1) // Hello

	for idx, runevalue := range s {
		fmt.Println("%#U starts at #d\n", runevalue, idx)
	}

	fmt.Println("\n Using DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i = i + w {
		runeValue, width := utf8.DecodeRuneInString(s[i:1])
		fmt.Println("%#U starts at %d\n", runeValue, i)
		w = width
		examineRune(runeValue)
	}

}

func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
