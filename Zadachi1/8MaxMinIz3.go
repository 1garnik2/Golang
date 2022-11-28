package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	fmt.Scanf("%v %v %v", &a, &b, &c)
	max := int(math.Max(math.Max(a, b), c))
	min := int(math.Min(math.Min(a, b), c))
	fmt.Println(max * max)
	fmt.Println(min * min)
}
