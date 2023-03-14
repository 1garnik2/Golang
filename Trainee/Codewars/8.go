// считает угол с
package main

import "fmt"

func main() {
	a, b := 45, 34
	fmt.Println(OtherAngle(a, b))
}

func OtherAngle(a int, b int) int {
	c := 180 - a - b
	return c
}
