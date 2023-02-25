package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, s float64
	fmt.Scan(&a, &b)
	s = (a * b) / 2
	c = math.Sqrt((math.Pow(a, 2)) + (math.Pow(b, 2)))
	fmt.Printf("Square =%.2f \n", s)
	fmt.Printf("Hipotenuse =%.2f", c)

}
