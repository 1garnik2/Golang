package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, d, x, x1, x2 float64
	fmt.Scan(&a, &b, &c)
	d = b*b - 4*a*c
	if d < 0 {
		fmt.Print("No solutions")
	} else {
		if d == 0 {
			x = -b / (2 * a)
			fmt.Printf("x = %f", x)
		} else {
			x1 = (-b - math.Sqrt(d)) / (2 * a)
			x2 = (-b + math.Sqrt(d)) / (2 * a)
			fmt.Printf(("x1 = %f", x1),("x2 = %f", x2))

		}
	}
}
