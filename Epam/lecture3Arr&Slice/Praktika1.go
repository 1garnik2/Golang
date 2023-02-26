package main

import "fmt"

// ввод 10 чисел для масива
func main() {
	var mas [100]int
	var n int
	fmt.Scan(&n)
	for i := 0; i < 100; i++ {
		fmt.Scan(&mas[i])
	}

	// сумма чисел масива
	sum := 0
	for i := 0; i < 100; i++ {
		if mas[i] > 0 {
			sum += 1
		}
	}
	fmt.Println(sum)
}
