package main

import "fmt"

var w = 20 // можно использовать за пределами функции

func main() {
	// x := 10  только в середине функции
	// y := 3.14987 только в середине функции

	var x int = 10
	var y float64
	y = 3.14987
	fmt.Printf("x = %d y = %10.4f \n", x, y)

}
