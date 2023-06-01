package main

import "fmt"

func MakeNegative(x int) int {
	if x > 0 {
		x = 0 - x
	}
	return 0 - x
}
func main() {
	fmt.Println(MakeNegative(-5))
}
