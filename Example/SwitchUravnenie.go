// тестовые данные 4 1 5     2 -4 2     2 -6 4
package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, d, x, x1, x2 float64
	fmt.Scan(&a, &b, &c)
	d = b*b - 4*a*c
	switch {
	case d < 0:
		fmt.Println("No solutions")
	case d == 0:
		x = -b / (2 * a)
		fmt.Printf("x = %f", x)
	default:
		x1 = (-b - math.Sqrt(d)/(2*a))
		x2 = (-b + math.Sqrt(d)/(2*a))
		fmt.Printf("x1 = %f", x1, "x2 = %f", x2)
	}
}
