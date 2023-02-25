package main

import "fmt"

func main() {
	slice := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	fmt.Println(slice)
	for _, i := range slice {
		fmt.Println(i)
	}
	sum := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			sum += slice[i][j]
		}
	}
	fmt.Println("Sum =", sum)
}
