package main

import "fmt"

func main() {
	//при изменении массива - измениться и слайс
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)

	// Создание слайса копирующего массив
	slice1 := arr[:]
	fmt.Println("slice1", slice1)

	//создание слайса с 0 по 4 инд. НЕ включительно
	slice2 := arr[:4]
	fmt.Println("slice2", slice2)

	//слайс с 1 инд до конца
	slice3 := arr[1:]
	fmt.Println("slice3", slice3)

	//слайс с 1 инд до 3. 3инд НЕ запишеться
	slice4 := arr[1:3]
	fmt.Println(slice4)
}
