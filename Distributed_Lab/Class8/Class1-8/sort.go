package main

import "fmt"

func main() {
	arr := []int{4, 3, 1, 5, 3, 14, 3, 4, 7, 10, 11, 12, 4, 3, 3, 56, 3, 12, 122, 3, 2, 3}
	var temp int                          // дополнительная переменная
	for i := 0; i < (len(arr) - 1); i++ { //выполнит 9 проходов по массиву
		for j := 0; j < (len(arr)-1)-i; j++ { // срав-е эл и уменьшение кол-ва проходов
			if arr[j] > arr[j+1] { /*сравн-е пар элементов
				  если первый больше чем второй*/
				temp = arr[j] /*в переменную записывается
				  индекс начиная 0 до 9*/
				arr[j] = arr[j+1] // увеличивает индекс на 1
				arr[j+1] = temp   // увеличиный индекс присваивается temp
			}
		}
	}
	fmt.Println(arr)
}