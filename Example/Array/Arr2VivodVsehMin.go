package main

import "fmt"

func main() {
	//ввод в двухмерный массив
	var arr [100][100]int
	var r, c, min int
	fmt.Scan(&r, &c)

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			fmt.Scan(&arr[i][j])
		}
	}
	//вывод мин значений всех строк

	for i := 0; i < r; i++ {
		min = arr[i][0]
		for j := 0; j < c; j++ {
			if arr[i][j] < min {
				min = arr[i][j]

			}
		}
		fmt.Print(min, " ")
	}

}
