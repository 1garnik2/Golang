// Напишите программу, принимающая на вход число N (N \geq 4)N(N≥4),
// а затем NN целых чисел-элементов среза.
// На вывод нужно подать 4-ый (3 по индексу) элемент данного среза.
package main

import "fmt"

func main() {
	var k int
	fmt.Scan(&k)
	slice := make([]int, k)
	for i := 0; i < k; i++ {
		fmt.Scan(&slice[i])
	}
	fmt.Println(slice[3])
}
