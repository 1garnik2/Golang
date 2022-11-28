package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int
	x = 5
	y = 4

	var m = int(math.Pow(float64(x), float64(y)))
	fmt.Println(m)
}
