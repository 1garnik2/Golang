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
	CylinderVolume() float64
	ConusVolume() float64
}

func (n Date) CylinderVolume() float64 {
	return Pi * (math.Pow(float64(n.Radius), 2)) * float64(n.High)
}
func (n Date) ConusVolume() float64 {
	return (Pi * math.Pow(float64(n.Radius), 2) / 3) * float64(n.High)
}
func main() {
	var a ValueCilindrConus
	numbers := Date{10, 6}
	a = numbers
	fmt.Printf("Cylinder volume %.d \n", int(a.CylinderVolume()))
	fmt.Printf("Conus volume %.d \n", int(a.ConusVolume()))
}
