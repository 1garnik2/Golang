package main

import "fmt"

func main() {
	// Создаем слайс со значениями
	slice := []string{"Vasya", "Petya", "Kate"}
	/*Создаем пустой слайс с длииной
	как у первого слайса*/
	slice1 := make([]string, len(slice))
	// Выполняем функцию copy
	copy(slice1, slice)

	fmt.Print(slice1)
}
