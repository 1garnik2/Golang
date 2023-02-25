package main

import "fmt"

func main() {
	//при изменении массива - измениться и слайс
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	i := 3
	// Создание слайса копирующего массив
	slice1 := arr[i:]
	fmt.Println("slice1", slice1)
}
