package main

import "fmt"

func main() {
	x := 2
	y := 3
	fmt.Println(x, y)
	f1(x, y)
	fmt.Println(x, y)
	f2(&x, &y)
}

func f1(x int, y int) {
	x = 5
	y = 6
	fmt.Println(x, y)
}
func f2(x *int, y *int) {
	*x = 5
	*y = 6
	fmt.Println(*x, *y)
}
