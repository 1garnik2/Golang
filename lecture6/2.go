package main

import "fmt"

func main() {
	arr1 := [2]int{1, 2}
	sum(arr1)
	fmt.Println(arr1)
	arr2 := []int{1, 2}
	sum2(arr2)
	fmt.Print(arr2)
}
func sum(arr [2]int) {
	s := 0
	arr[0] = 100
	for _, x := range arr {
		s += x
	}
	fmt.Println(arr, "sum =", s)
}
func sum2(arr []int) {
	s := 0
	arr[0] = 100
	for _, x := range arr {
		s += x
	}
	fmt.Println(arr, "sum =", s)
}
