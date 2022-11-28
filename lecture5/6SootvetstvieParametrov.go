package main

import "fmt"

func main() {
	x := 2
	y := 3
	sum1(x, y)
	s := sum2(x, y)
	fmt.Println(s)

	fmt.Println(sum2(x, y))

	sum1(sum2(3, 2), sum2(4, 4))

	a := 1.5
	//b:=3.14
	sum3(x, a) // сначала инт потом флоат!
	//sum3(b,y)
	//sum3(x,y)

}
func sum1(x int, y int) {
	fmt.Println(x + y)
}
func sum2(x int, y int) int {
	return x + y
}
func sum3(x int, y float64) {
	fmt.Println(float64(x) + y)
}
