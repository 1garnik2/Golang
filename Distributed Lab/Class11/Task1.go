package main

import "fmt"

type Date struct {
	Square int
	High   int
}
type ValueInteface interface {
	PrismVolume() int
	PyramidVolume() float64
}

func (n Date) PrismVolume() int {
	return n.Square * n.High
}
func (n Date) PyramidVolume() float64 {
	return (float64(n.Square) / 3) * float64(n.High)
}
func main() {
	var a ValueInteface
	numbers := Date{Square: 12, High: 6}
	a = numbers
	fmt.Printf("Prism volume %d \n", a.PrismVolume())
	fmt.Printf("Pyramid volume %.1f \n", a.PyramidVolume())
}
