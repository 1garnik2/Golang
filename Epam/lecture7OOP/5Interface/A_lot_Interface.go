package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimetr() float64
}

// struct
type Rectangle struct {
	length, // type float64
	heigth float64
}

// =======================================
// метод для структуры Rectangle с методом area() площадь
// возвращает float64
func (r Rectangle) area() float64 {
	return r.heigth * r.length
}

// метод для структуры Rectangle с методом perimetr() периметр
// возвращает float64
func (r Rectangle) perimetr() float64 {
	return (r.length + r.heigth) * 2
}

// ==========================================

type Triangle struct {
	sideA, sideB, sideC float64
}

func (t Triangle) area() float64 {
	p := (t.sideA + t.sideB + t.sideC) / 2
	return math.Sqrt(p * (p - t.sideA) * (p - t.sideB) * (p - t.sideC))
}

func (t Triangle) perimetr() float64 {
	return t.sideA + t.sideB + t.sideC
}

// вычисляет площадь любой из фигур
func calculate(s Shape) (float64, float64) {
	return s.area(), s.perimetr()
}
func main() {
	// создание переменных с типом структур и
	// заполнение структур значениями
	r := Rectangle{4, 8}        // прямоуголник
	t := Triangle{4, 7, 9}      // треугольник
	area, sqwer := calculate(r) //возвращает 2 переменные
	fmt.Printf("Площадь %.2f Периметр %.2f\n", area, sqwer)

	fmt.Println(calculate(t))

}
