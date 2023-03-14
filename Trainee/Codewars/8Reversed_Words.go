package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "The greatest victory is that which requires no battle"
	fmt.Println(ReverseWords(str))
}

func ReverseWords(str string) string {
	arr := strings.Split(str, " ")
	revers := ""
	for _, v := range arr {
		revers = v + " " + revers
	}

	return strings.TrimSuffix(revers, " ") // reverse those words
}
