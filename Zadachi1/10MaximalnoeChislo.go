package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	var i int
	fmt.Scanf("%d", &i)
	a = float64(i / 100)
	b = float64(i / 10 % 10)
	c = float64(i % 10)
	max := int(math.Max(math.Max(a, b), c))
	min := int(math.Min(math.Min(a, b), c))
	mid := int(a) + int(b) + int(c) - max - min
	fmt.Println(max*100 + mid*10 + min)

}
