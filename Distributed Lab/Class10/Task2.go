package main

import (
	"fmt"
	"math"
)

func main() {
	var AC, BC, ACB float64
	AC = 22
	BC = 21
	ACB = 60
	AB := math.Pow(AC, 2) + math.Pow(BC, 2) - 2*AC*BC*(math.Cos(ACB))
	fmt.Print(math.Round(AB))
}
