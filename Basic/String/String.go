package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "的期望带我去带我去"
	fmt.Printf("%X \n", s)
	for _, b := range []byte(s) {
		fmt.Printf("%X \n", b)
	}

	for i, ch := range s { // ch is a rune
		fmt.Printf("%d %X \n", i, ch)
	}

	fmt.Println("rune count:", utf8.RuneCountInString(s))

	bytes := []byte(s)
	ch, size := utf8.DecodeRune(bytes)
	fmt.Printf("%c", ch)
	fmt.Println()
	fmt.Println(size) // 中文size=3 英文1

	for i, ch := range []rune(s) {
		fmt.Printf("%d : %c", i, ch)
		fmt.Println()
	}
}
