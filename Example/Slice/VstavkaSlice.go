package main

import "fmt"

func main() {
	//инициализация слайса
	slice := []int{1, 2, 3, 4, 5}

	fmt.Println(slice)
	//добавление элемента в середину слайса
	slice = append(slice, 0)

	copy(slice[3:], slice[2:])
	slice[2] = 7
	fmt.Println(slice)
}
