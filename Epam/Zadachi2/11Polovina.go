package main

import "fmt"

func main() {
	var arr [100][100]int
	var n int
	fmt.Scan(&n)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			arr[i][j] = 2
			arr[n-1-j][n-1-i] = 1

		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}
}
