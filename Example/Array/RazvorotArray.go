package main

import "fmt"

// ввод 11 чисел для масива
func main() {
	var arr = [11]int{1, 2, 3, 4, 5, 14, 7, 8, 9, 10, 11}

	fmt.Println(arr)

	// разворот массива (перезапись)
	var temp int
	for i := 0; i < 6; i++ { //кол-во пар!!!!
		temp = arr[i]
		arr[i] = arr[10-i]
		arr[10-i] = temp
	}
	fmt.Println(arr)
}
