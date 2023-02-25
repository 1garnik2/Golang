package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "Hello, today is January 11. I am 28 now. Date of birth 11.01.1990."

	fmt.Println("Test string", str)

	re := regexp.MustCompile("[0-9]+")
	fmt.Println(re.FindAllString(str, -1))
}
