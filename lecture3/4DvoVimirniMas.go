package main

import "fmt"

func main() {
	mas := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(mas)
	fmt.Println(mas[1][1])

	for i := 0; i < 3; i++ { //перебириает рядки
		fmt.Println(mas[i]) //и выводит с нового рядка
	}

	for i := range mas {
		fmt.Println(mas[i])
	}
	for i := range mas {
		fmt.Println(i)
	}
}
