package main

import "fmt"

func main() {
	//заполнение массива с клавиатуры
	var arr [10]int
	for i := 0; i < 10; i++ {
		fmt.Scan(&arr[i])
	}
	// максимальное значение
	//Обращение к элементу за индексом
	index := 0
	for i := 1; i < 10; i++ {
		if arr[i] > arr[index] {
			index = i
		}
	}
	fmt.Println(index)              // вывод индекса максимального числа
	fmt.Println("max=", arr[index]) //вывод значения индекса
}
