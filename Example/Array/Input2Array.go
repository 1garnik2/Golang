package main

import "fmt"

func main() {
	var arr [3][4]int
	// заполнение массива
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			fmt.Scan(&arr[i][j])
		}
	}
	//вывод массива
	for _, i := range arr {
		fmt.Println(i)
	}
}
