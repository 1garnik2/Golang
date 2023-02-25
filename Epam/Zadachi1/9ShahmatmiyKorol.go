package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64
	var n int
	fmt.Scanf("%d", &n)
	var x, y, m, t int
	fmt.Scanf("%d %d %d %d", &x, &y, &m, &t)
	a = math.Abs(float64(x - m))
	b = math.Abs(float64(y - t))
	fmt.Println(int(math.Max(a, b)))
}
