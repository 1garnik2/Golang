package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}
type circl struct {
	radius float64
}

type squar struct {
	a float64
	b float64
}

type triang struct {
	a float64
	b float64
	c float64
}

func (t triang) Area() float64 { //по формуле Герона
	p := (t.a + t.b + t.c) / 2
	S := math.Sqrt(p * (p - t.a) * (p - t.b) * (p - t.c))
	return S
}

func (c circl) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}
func (s squar) Area() float64 {
	return s.a * s.b
}

func main() {
	var circle Shape = circl{6.45}
	var triangle Shape = triang{4, 5.6, 8}
	var square Shape = squar{4, 6}
	fmt.Printf("Area square = %.2f\n", square.Area())
	fmt.Printf("Area triangle = %.2f\n", triangle.Area())
	fmt.Printf("Area circle = %.2f\n", circle.Area())
}
