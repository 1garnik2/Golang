package main

import "fmt"

func main() {
	//заполнение двумерного массива
	var arr [100][100]int
	var n, m int
	fmt.Scan(&n, &m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scan(&arr[i][j])
		}
	}
	min := arr[0][0]
	nmin := 0 //порядковый номер рядка
	max := arr[0][0]
	nmax := 0

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			//минимальный
			if min > arr[i][j] {
				min = arr[i][j]
				nmin = i //индекс рядка
			}
			//максимальный рядок
			if max <= arr[i][j] {
				max = arr[i][j]
				nmax = i //индекс рядка
			}
		}
	}
	//поменять местами одновимимрные массивы
	arr[nmax], arr[nmin] = arr[nmin], arr[nmax]

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Println(arr[i][j], " ")
		}
	}
}
