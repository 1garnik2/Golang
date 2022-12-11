package main

import (
	"fmt"
	"math"
)

const Pi = 3.14

type Date struct {
	Radius int
	High   int
}
type ValueCilindrConus interface {
	СylinderVolume() float64
	ConusVolume() float64
}

func (n Date) СylinderVolume() float64 {
	return Pi * (math.Pow(float64(n.Radius), 2)) * float64(n.High)
}
func (n Date) ConusVolume() float64 {
	return (Pi * math.Pow(float64(n.Radius), 2) / 3) * float64(n.High)
}
func main() {
	var a ValueCilindrConus
	numbers := Date{10, 6}
	a = numbers
	fmt.Printf("Сylinder volume %.d \n", int(a.СylinderVolume()))
	fmt.Printf("Conus volume %.d \n", int(a.ConusVolume()))
}
