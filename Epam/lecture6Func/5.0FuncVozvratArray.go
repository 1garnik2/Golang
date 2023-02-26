package main

import "fmt"

func main() {
	arr1 := [3]int{3, 4, 5}
	fmt.Println(mul2(arr1))
	arr2 := [3]int{1, 2, 3}
	fmt.Println(mul2(arr2))
}

//Фунуция умножения элементов массива на 2
func mul2(arr [3]int) [3]int {
	var rez [3]int
	for i := range arr {
		rez[i] = arr[i] * 2
	}
	return rez
}
