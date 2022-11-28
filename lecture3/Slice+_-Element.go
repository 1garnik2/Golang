package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}

	fmt.Println(slice)
	// Добавление элемента в конец слайса
	slice = append(slice, 6)
	fmt.Println(slice)
	/* Удаление Элемента
	дабовляем новый слайс"slice[0:3]" и продолжаем
	с 3 позиции "slice[3:]..."*/
	slice = append(slice[0:2], slice[3:]...)
	fmt.Println(slice)
	/* Удаление Элемента
	смещение слайса на одну позицию влево*/
	copy(slice[1:], slice[2:])
	slice = slice[:len(slice)-1] // последние элементы одинаковые по этому последний del
	fmt.Println(slice)

	// добавление в слайс
	slice = append(slice, 0)
	copy(slice[3:], slice[2:])
	slice[2] = 7
	fmt.Println(slice)
}
