package main

import "fmt"

func main() {
	//Создаем новый слайс 1 емкостью 3
	// на 2 элемента
	slice := make([]string, 2, 3)

	//Слайс со значениями
	source := []string{"Vasya", "Petya", "Kate"}

	//Функция копирования из одного
	//слайса в другой. Функция copy
	//возвращает кол-во скопированных элементов
	copySlice := copy(slice, source)

	//Вывод скопированного слайса
	fmt.Println(slice)

	//Вывод количества скопированных элементов
	fmt.Print(copySlice)
}
