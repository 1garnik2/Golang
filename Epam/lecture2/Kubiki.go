package main

import "fmt"

func main() {
	var a, b, x, y, s int
	fmt.Scanf("%d %d %d %d", &a, &b, &x, &y)
	s = (a + b + x + y)
	switch s {
	case 14:
		fmt.Println(1)
	case 13:
		fmt.Println(2)
	case 12:
		fmt.Println(3)
	case 11:
		fmt.Println(4)
	case 10:
		fmt.Println(5)

	}
}
