package main

import (
	"fmt"
	"math"
)

func main() {
	var arr [100]int
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	//сортировка пузірьковаяпорядок возростания модулей
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if (math.Abs(float64(arr[j])) > math.Abs(float64(arr[j+1]))) ||
				(math.Abs(float64(arr[j])) == math.Abs(float64(arr[j+1])) && arr[j+1] < 0) {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
}
