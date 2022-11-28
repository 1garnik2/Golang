// для того что б посчиать символы используют rune
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	var str1 = "bravo"
	fmt.Println(str1, len(str1))

	var str2 = "браво"
	fmt.Println(str2, len(str2))

	fmt.Println(len(str1), utf8.RuneCountInString(str1))
	fmt.Println(len(str2), utf8.RuneCountInString(str2))
}
