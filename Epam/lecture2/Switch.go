package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)

	switch a {
	case 0:
		fmt.Println("a = 0")
	case 1:
		fmt.Println("a = 1")
	case 4:
		fmt.Println("a = 2")
	case 5:
		fmt.Println("a = 3")
	case 7:
		fmt.Println("a = 4")
	default:
		fmt.Println("NO")
	}
}
