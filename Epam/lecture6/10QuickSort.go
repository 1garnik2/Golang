package main

import "fmt"

func main() {
	//Слайс который нужно отсортировать
	arr := []int{5, 3, 8, 2, 1, 9, 4, 10, 7, 6}
	qsort(arr, 0, 9)
	fmt.Println(arr)
}
func qsort(arr []int, first int, last int) {
	if first < last {
		left := first
		right := last
		/*Выбор опорного элемента
		(выбираем середину массива)*/
		middle := arr[(left+right)/2]
		/* Перенос всех элементов которые
		меньше опортного в левую часть*/
		for left <= right {
			for arr[left] < middle {
				left++
			}
			for arr[right] > middle {
				right--
			}
			if left <= right {
				arr[left], arr[right] = arr[right], arr[left]
				left++
				right--
			}
		}
		qsort(arr, first, right)
		qsort(arr, left, last)
	}
}
