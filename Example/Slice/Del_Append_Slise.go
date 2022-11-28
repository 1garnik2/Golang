package main

import "fmt"

func main() {
	//инициализация слайса
	slice := []int{1, 2, 3, 4, 5}

	fmt.Println(slice)
	//добавление элемента в конец слайса
	slice = append(slice, 6)
	fmt.Println(slice)
	//удаление элемента
	slice = append(slice[0:2], slice[3:]...)
	fmt.Println(slice)
	//удаление элемента
	copy(slice[1:], slice[2:])
	slice = slice[:len(slice)-1]
	fmt.Println(slice)
	//
}
