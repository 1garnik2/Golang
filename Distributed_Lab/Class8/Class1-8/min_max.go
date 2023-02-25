package main

import "fmt"

func main() {
	arr := []int{4, 3, 1, 5, 3, 14, 3, 4, 7, 10, 11, 12, 4, 3, 3, 56, 3, 12, 122, 3, 2, 3}

	min := arr[0]
	nmin := 0 //порядковый номер рядка
	max := arr[0]
	nmax := 0

	for i := 0; i < len(arr); i++ {

		//минимальный
		if min > arr[i] {
			min = arr[i]
			nmin = i //индекс рядка
		}
		//максимальный рядок
		if max <= arr[i] {
			max = arr[i]
			nmax = i //индекс рядка
		}
	}

	//поменять местами одновимимрные массивы
	arr[nmax], arr[nmin] = arr[nmin], arr[nmax]

	for i := 0; i < len(arr); i++ {

		fmt.Print(arr[i], " ")
	}

}
