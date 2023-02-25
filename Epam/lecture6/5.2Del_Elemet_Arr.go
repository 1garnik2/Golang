// Получает массив из 3-х элементов
// а возвращает из 2-х
package main

import "fmt"

func main() {
	arr1 := [3]int{3, 4, 5}
	//Удаляет по индеску
	fmt.Println(deleteElement(arr1, 1))

}

func deleteElement(arr [3]int, index int) [2]int {
	//Создаем новый массив
	var rez [2]int
	j := 0
	for i := 0; i < 3; i++ {
		if i == index {
			continue
		}
		rez[j] = arr[i]
		j++
	}
	return rez
}
