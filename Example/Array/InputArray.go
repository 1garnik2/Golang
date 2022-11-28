package main

import "fmt"

// ввод 10 чисел для масива
func main() {
	var arr [10]int
	for i := 0; i < 10; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Println(arr)
}
