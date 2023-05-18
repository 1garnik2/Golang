package main

import "fmt"

type Number interface {
	int | float64
}

func main() {
	arr1 := []int{1, 2, 3, 4, 5, 4, 3, 2}
	arr2 := []float64{1, 0, -2, 3, 4, 5, 4, 3, 2}
	fmt.Println("int: ", Min(arr1))
	fmt.Println("float64: ", Min(arr2))
}
func Min[T Number](arr []T) T {
	var sl = arr[0]
	for _, val := range arr {
		if val < sl {
			sl = val
		}
	}
	return sl
}
