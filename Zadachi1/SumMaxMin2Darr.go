package main

import "fmt"

func main() {
	//запись в массив
	var arr [100][100]int
	var r, s int
	fmt.Scan(&r, &s)

	for i := 0; i < r; i++ { //перебирает № рядка
		for j := 0; j < s; j++ { //перебирает № столбиков
			fmt.Scan(&arr[i][j])
		}
	}
	//сумма масива
	sum := 0
	for _, row := range arr {
		for _, x := range row {
			sum += x
		}
	}
	//max значение
	max := arr[0][0]
	for i := 0; i < r; i++ {
		for j := 0; j < s; j++ {
			if arr[i][j] > max {
				max = arr[i][j]

			}
		}
	}
	//min значение
	min := arr[0][0]
	for i := 0; i < r; i++ {
		for j := 0; j < s; j++ {
			if arr[i][j] < min {
				min = arr[i][j]

			}
		}
	}

	fmt.Println(sum, min, max)
}
