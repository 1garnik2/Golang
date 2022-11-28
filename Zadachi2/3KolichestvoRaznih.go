package main

import (
	"fmt"
)

func main() {
	var arr [100]int
	var n, count int
	fmt.Scan(n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	for i := 0; i < n-1; i++ {
		if arr[i] != arr[+1] {
			count++
		}
	}
	fmt.Println(count + 1)
}
