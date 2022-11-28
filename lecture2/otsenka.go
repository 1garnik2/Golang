package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	switch {
	case a < 60:
		fmt.Println("rating unsatisfactory")
	case a > 60 && a < 74:
		fmt.Println("rating satisfactory")
	case a >= 74 && a < 90:
		if a == 77 {
			fmt.Println("Wow! Two sevens!")
			//break
		}
		fmt.Println("rating good")
	case a > 90 && a <= 100:
		fmt.Println("rating excellent")
	default:
		fmt.Println("No such rating exist")

	}
}
