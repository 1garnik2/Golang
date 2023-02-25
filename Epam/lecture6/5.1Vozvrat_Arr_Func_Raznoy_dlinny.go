// Получает массив из 3-х элементов
// а возвращает из 4-х
package main

import "fmt"

func main() {
	arr1 := [3]int{3, 4, 5}
	//добавляет элемент в конец массива
	fmt.Println(appendElement(arr1, 6))
}
func appendElement(arr [3]int, item int) [4]int {
	//создаем новый пустой массив
	var rez [4]int
	// первые 3 элемента копируем
	for i := range arr {
		rez[i] = arr[i]
	}
	// в индекс последнего добавляем элемент
	rez[3] = item
	return rez
}
