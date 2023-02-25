package main

import (
	"fmt"
	"strconv"
)

func main() {

	var arr = [10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Print(CreatePhoneNumber(arr)) // returns "(123) 456-7890"

}

func CreatePhoneNumber(numbers [10]uint) string {
	//var format = "(xxx) xxx-xxxx"

	var str, tel string
	for i := 0; i < 10; i++ {
		str += strconv.Itoa(int(numbers[i]))
	}
	//strings.Split(str, "")
	kod := str[0:3]
	per := str[3:6]
	las := str[6:]
	tel = "(" + kod + ") " + per + "-" + las
	return tel

}
