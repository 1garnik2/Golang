package main

import "fmt"

func main() {
	slice := [][]int{{1, 2, 3, 4}, {5, 6}, {7, 8, 9}}

	fmt.Println(slice)
	for _, i := range slice {
		fmt.Println(i)
	}
	sum := 0
	for _, row := range slice {
		for _, col := range row {
			sum += col
		}
	}
	fmt.Println("Sum =", sum)
}
