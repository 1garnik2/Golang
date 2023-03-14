package main

import (
	"fmt"
	"unicode"
)

func troo_or_false(pass string) bool {
	if len(pass) < 5 {
		return false
	}
	srezrun := []rune(pass)
	for _, simvol := range srezrun {
		if !unicode.Is(unicode.Latin, simvol) && !unicode.Is(unicode.Digit, simvol) {
			return false
		}
	}

	return true
}
func main() {
	var text string
	fmt.Scan(&text)
	if troo_or_false(text) {
		fmt.Print("Ok")
	} else {
		fmt.Print("Wrong password")
	}
}
