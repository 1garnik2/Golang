package main

import (
	"fmt"
	"math"
)

func main() {
	var age = mars_age(1000)
	fmt.Println(age) // 9

}

func mars_age(x int) int {

	s := math.Sqrt(365 * x / 687)
	return s

}
