package main

import "fmt"

func main() {
	x := 2
	y := 3
	fmt.Println(x, y) //значения х у == 2 3
	f1(x, y)          //Только в функции меняет значение!
	fmt.Println(x, y) //Значения х у не изменились
	f2(&x, &y)        // предача х у в фунцию по ссылке
	fmt.Println(x, y) // Значения х у изменились
}

// предача КОПИЙ значений
func f1(x int, y int) {
	x = 5
	y = 6
	fmt.Println(x, y)
}

// ПЕРЕДАЧА ПО ЗНАЧЕНИЮ
func f2(x *int, y *int) {
	*x = 5
	*y = 6
	fmt.Println(*x, *y)
}
