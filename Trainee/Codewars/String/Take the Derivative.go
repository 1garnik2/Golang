package main

import (
	"fmt"
	"strconv"
)

func Derive(coefficient, exponent int) string {
	first := coefficient * exponent
	second := 0
	if exponent > 2 {
		second += exponent - 1
	} else {
		second = exponent
	}
	stroka := strconv.Itoa(first) + "x^" + strconv.Itoa(second)
	return stroka
}

func main() {
	fmt.Println(Derive(5, 3))

}
