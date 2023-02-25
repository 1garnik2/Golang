package main // второе число ставит на первое место

import "fmt"

func main() {
	var a, b, c, d int32
	fmt.Scan(&a)
	b = (a % 100) / 10
	c = a % 10
	d = a / 100
	fmt.Println((b * 100) + (c * 10) + d)
}
