/*
Задача 3.
Створити структуру, що містить поля - координати точки на площині.
За допомогою цієї структури визначити 3 точки трикутника (2,-4), (-5,-6), (1,3).
Обчислити площу трикутника.
*/
package main

import (
	"fmt"
	"math"
)

type Points struct {
	x float64
	y float64
}

func main() {
	a := Points{2, -4}
	b := Points{-5, -6}
	c := Points{1, 3}
	ab := math.Sqrt(math.Pow((b.x-a.x), 2) + math.Pow((b.y-a.y), 2))
	bc := math.Sqrt(math.Pow((c.x-b.x), 2) + math.Pow((c.y-b.y), 2))
	ca := math.Sqrt(math.Pow((a.x-c.x), 2) + math.Pow((a.y-c.y), 2))
	p := (ab + bc + ca) / 2
	S := math.Sqrt(p * (p - ab) * (p - bc) * (p - ca))
	fmt.Printf("%.2f", S)
}
