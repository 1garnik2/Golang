package main

import "fmt"

func main() {
	x := 2
	y := 3
	fmt.Println(sum(x, y))

	var area, perimetr = culc(x, y)
	fmt.Println("Area =", area)
	fmt.Println("Perimetr =", perimetr)
}

//возвращает 1 результат
func sum(x int, y int) (z int) {
	z = x + y
	return
}

//возвращает 2 результата
func culc(a int, b int) (int, int) {
	var s = a * b
	var p = (a + b) * 2
	return s, p
}

// abo tac
func culc2(a int, b int) (s int, p int) {
	s = a * b
	p = (a + b) * 2
	return
}
