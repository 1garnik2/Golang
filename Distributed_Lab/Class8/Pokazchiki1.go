package main

import "fmt"

func swap(x *int, y *int) {
	*x, *y = *y, *x
}
func main() {
	x := 2
	y := 1
	swap(&x, &y)
	fmt.Printf("x = %d \ny = %d", x, y)
}
