package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	n = int(math.Abs(float64(n)))
	if n == 0 {
		fmt.Println(1)
	} else {
		s := 0
		for n > 0 {
			if n%10 == 0 {
				s++
			}
		}
	}
}
