// Створити з нього зріз, що містить тільки парні елементи масиву
//
//	і відсортувати його за зростанням, вивести зріз.
package main

import "fmt"

func main() {
	arr := [...]int{48, 96, 86, 68, 57, 82, 63, 70}
	var Slice []int
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 0 {
			Slice = append(Slice, arr[i])
		}
	}
	for i := 0; i < len(Slice); i++ {
		for j := i; j < len(Slice); j++ {
			if Slice[i] > Slice[j] {
				tmp := Slice[i]
				Slice[i] = Slice[j]
				Slice[j] = tmp
			}
		}
	}
	fmt.Println(Slice)
}
