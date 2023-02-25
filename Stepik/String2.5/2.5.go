package main

import (
	"fmt"
	"strings"
)

/*

Sample Input:
ihgewlqlkot

Sample Output:
hello
*/

func main() {
	var text string
	_, _ = fmt.Scan(&text)

	var b strings.Builder
	for i, r := range text {
		if i%2 != 0 {
			b.WriteRune(r)
		}
	}

	fmt.Println(b.String())
}
