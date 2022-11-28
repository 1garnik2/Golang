package main

import "fmt"

func main() {
	arr1 := [3]int{3, 4, 5}
	fmt.Println(mul2(arr1))
	arr2 := [3]int{1, 2, 3}
	fmt.Println(mul2(arr2))

	fmt.Println(appendElement(arr1, 6))
	fmt.Println(deleteElement(arr2, 1))
}

func mul2(arr [3]int) [3]int {
	var rez [3]int
	for i := range arr {
		rez[i] = arr[i] * 2
	}
	return rez
}
func appendElement(arr [3]int, item int) [4]int {
	var rez [4]int
	for i := range arr {
		rez[i] = arr[i]
	}
	rez[3] = item
	return rez
}
func deleteElement(arr [3]int, index int) [2]int {
	var rez [2]int
	j := 0
	for i := 0; i < 3; i++ {
		if i == index {
			continue
		}
		rez[j] = arr[i]
		j++
	}
	return rez
}
