package main

import "fmt"

func main() {
	n, x, y := 12, 3, 4
	fmt.Println(IsDivisible(n, x, y))
}
func IsDivisible(n, x, y int) bool {
	// your code here
	if n%x == 0 && n%y == 0 {
		return true
	}
	return false
}
