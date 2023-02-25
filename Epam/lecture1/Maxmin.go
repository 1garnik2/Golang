package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y float64
	x = 3.14
	y = 4.5
	var m float64

	m = math.Max(x, y)
	fmt.Println("Max =", m)

	m = math.Min(x, y)
	fmt.Println("Min =", m)

	var c, v int
	c = 3
	v = 4
	var b int

	b = int(math.Min(float64(c), float64(v)))
	fmt.Println("cv", b)
}
