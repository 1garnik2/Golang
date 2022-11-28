package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	n = int(math.Abs(float64(n)))
	s := 0
	for n > 0 {
		s += n % 10
		n /= 10
	}
	fmt.Println(s)
}
