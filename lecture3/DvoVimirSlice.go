package main

import "fmt"

func main() {
	slice1 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	fmt.Println("slice1")
	for _, i := range slice1 {
		fmt.Println(i)
	}
	sum := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			sum += slice1[i][j]
		}
	}
	fmt.Println("sum =", sum)

	fmt.Println()
	slice2 := [][]int{{1, 2, 3, 4}, {5, 6}, {7, 8, 9}}
	fmt.Println("slice2")
	for _, i := range slice2 {
		fmt.Println(i)
	}
	sum = 0
	for _, row := range slice2 {
		for _, x := range row {
			sum += x
		}
	}
	fmt.Println("sum =", sum)
}
