package main

import (
	"fmt"
	"unicode"
)

func main() {
	//text := "The30quick20brown10f0x1203jumps914ov3r1349the102l4zy dog"
	proba := "The30qu"
	SumOfIntegersInString(proba)

}
func SumOfIntegersInString(strng string) int {

	//re := regexp.MustCompile("[0-9]+")
	//fmt.Println(re.FindAllString(strng, -1))
	//return -1
	var s rune
	s = ' '
	srezrun := []rune(strng)
	for _, simvol := range srezrun {
		if !unicode.Is(unicode.Digit, simvol) {
			simvol = s
		}

	}
	fmt.Println(string(srezrun))
	return -1
}
