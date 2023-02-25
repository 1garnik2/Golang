package main

import "fmt"

func main() {
	var arr [3][3]int
	// заполнение массива
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Scan(&arr[i][j])
		}
	}
	//вывод массива
	for _, i := range arr {
		fmt.Println(i)
	}

	//сумма всех чисел главной диагонали
	sumd := 0
	for i := 0; i < 3; i++ {
		sumd += arr[i][i]

	}
	fmt.Println("sum diag = ", sumd)
	// сумма всех чисел побочной диагонали
	sumpd := 0
	for i := 0; i < 3; i++ {
		sumpd += arr[i][2-i]

	}
	fmt.Println("sum diag = ", sumpd)

}
