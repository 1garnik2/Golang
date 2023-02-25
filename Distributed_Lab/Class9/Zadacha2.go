/*
Задача 2.
Створити структуру, що містить поля - координати точки на площині.
За допомогою цієї структури визначити 3 точки прямокутника (1,2), (1,4),
(7,4).  Знайти площу прямокутника
*/
package main

import (
	"fmt"
)

type Point struct {
	x, y int
}

func main() {
	a := Point{1, 2}
	b := Point{1, 4}
	c := Point{7, 4}
	ab := b.y - a.y
	bc := c.x - b.x
	fmt.Println(ab * bc)
}
