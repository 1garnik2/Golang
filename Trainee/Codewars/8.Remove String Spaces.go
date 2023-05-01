package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "8aaaaa dddd r    g f"
	fmt.Println(NoSpace(text))
}
func NoSpace(word string) string {
	Newword := strings.ReplaceAll(word, " ", "")
	return Newword
}
