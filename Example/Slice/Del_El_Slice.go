package main

import "fmt"

func main() {
	//Создаем слайс
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// Индекс который надо удалить
	i := 3
	// удаление через append
	del_el := append(slice[:i], slice[i+1:]...)
	// Слайс 1 ПОВРЕЖДЕН
	fmt.Print(slice)
	fmt.Print(del_el)

}
