package main

import (
	"fmt"
	"strings"
)

func ModifySpaces(s, mode string) string {
	switch {
	case mode == "dash":
		s = strings.ReplaceAll(s, " ", "-")
	case mode == "underscore":
		s = strings.ReplaceAll(s, " ", "_")
	case mode == "unknown":
		s = strings.ReplaceAll(s, " ", "*")
	default:
		s = strings.ReplaceAll(s, " ", "*")
	}
	return s
}
func main() {
	fmt.Println(ModifySpaces(" hello world ", "dash"))
	fmt.Println(ModifySpaces(" how r u doing ", "underscore"))
	fmt.Println(ModifySpaces(" great ", ""))
	fmt.Println(ModifySpaces(" nice ", "unknown"))
}
