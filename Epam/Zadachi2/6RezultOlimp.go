package main

import "fmt"

func main() {
	//ввод для двух массивов
	var arr, num [100]int
	var n int
	fmt.Scan(&n)
	// первый массив заполняем данными с кол-м очков
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
		//второй массив заполняем порядковым номером учасников
		num[i] = i + 1
	}
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			//сортируем первый массив пузырьковым сортированием по убыванию
			//одновременно сортируем второй массив с порядковым номером
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				num[j], num[j+1] = num[j+1], num[j]
			}
		}
	}
	for i := 0; i < n; i++ {
		fmt.Print(num[i], " ")
	}
}
