package main

import (
	"fmt"
)

func main() {
	var world = "Test"

	fmt.Println(Solution(world))
}

func Solution(word string) string {
	var result = ""
	for _, c := range word {
		result = string(c) + result
	}
	return result
}
