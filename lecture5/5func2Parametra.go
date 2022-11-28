package main

import "fmt"

func main() {
	x := 2
	y := 4
	sum1(x, y)
	a := sum2(x, y)
	fmt.Println(a)
	fmt.Println(sum2(x, y))

	sum1(sum2(3, 2), sum2(2, 6))
}
func sum1(x int, y int) {
	fmt.Println(x + y)
}
func sum2(x int, y int) int {
	return x * y
}
