package main

import "fmt"

func main() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} //массив
	var slice []int                                //слайс
	slice = arr[1:4]                               // заполнение слайса с 1-го по 4-й НЕ включительно!

	fmt.Println(slice)
	for _, i := range slice {
		fmt.Print(i, " ")
	}
	fmt.Println()
	//изменение 0-го элемента слайса и 1-го элемента массива
	slice[0] = 6
	fmt.Println("array", arr)
	for _, i := range arr {
		fmt.Print(i, " ")
	}
	fmt.Println()
	//изменение 2-го элемента массива
	arr[2] = 15
	fmt.Print("Print slice:")
	for _, i := range slice {
		fmt.Print(i, " ")
	}
	fmt.Println()
	// добавление єлемента в слайс
	slice = append(slice, 16)
	fmt.Print("Print Slice:")
	for _, i := range slice {
		fmt.Print(i, " ")
	}
	fmt.Println()
	//выводим массив
	fmt.Print("array:")
	for _, i := range arr {
		fmt.Print(i, " ")
	}
	fmt.Println()
	//слайс2 инициализация от слайса
	var slice2 []int
	slice2 = slice[1:3]
	fmt.Print("slice2:")
	for _, i := range slice2 {
		fmt.Print(i, " ")
	}
	fmt.Println()

	slice2[1] = 20 //изменение 1 значения слайс2
	fmt.Println(arr)
	fmt.Println(slice)
	fmt.Println(slice2)

}
