package main

import (
	"fmt"
	"math"
)

const Pi = 3.14

type KoloShar struct {
	Radius int
}
type Area interface {
	koloArea() float64
	SharArea() float64
}

func (n KoloShar) koloArea() float64 {
	return math.Pow(float64(n.Radius), 2) * Pi
}
func (n KoloShar) SharArea() float64 {
	return 4 * Pi * math.Pow(float64(n.Radius), 2)
}
func main() {
	var s Area
	dannye := KoloShar{10}
	s = dannye
	fmt.Printf("Area Kola %.d \n", int(s.koloArea()))
	fmt.Printf("Area Shar %.d \n", int(s.SharArea()))

}