package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "abs"
	b := 2
	fmt.Println(RepeatStr(a, b))
}
func RepeatStr(n string, s int) string {
	return strings.Repeat(n, s)
}
