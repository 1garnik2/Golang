package main

import "fmt"

// ввод 10 чисел для масива
func main() {
	var arr = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(arr)

	// максимальное значение
	index := 0
	for i := 1; i < 10; i++ {
		if arr[i] > arr[index] {
			index = i
		}
	}
	fmt.Println("max=", arr[index])

}
