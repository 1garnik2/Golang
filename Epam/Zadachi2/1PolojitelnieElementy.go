package main

import "fmt"

func main() {
	var arr [100]int
	var n, count int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
		if arr[i] > 0 {
			count++
		}
	}
	fmt.Println(count)
}

/*
package main

import "fmt"

// ввод 100 чисел для масива
func main() {
	var mas [100]int
	var n int
	fmt.Scan(&n)
	for i := 0; i < 100; i++ {
		fmt.Scan(&mas[i])
	}

	// сумма положительніх чисел масива
	sum := 0
	for i := 0; i < 100; i++ {
		if mas[i] > 0 {
			sum += 1
		}
	}
	fmt.Println(sum)
}
*/
