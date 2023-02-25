package main

import "fmt"

func main() {
	//запись в массив
	var arr [100]int
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ { //ввод значений в массив
		fmt.Scan(&arr[i])
	}
	//первое минимальное значение
	min := 0
	for i := 0; i < n; i++ {
		if arr[i] < arr[min] {
			min = i
		}
	}
	//fmt.Println("индекс min", min)
	//первое четное значение
	l := 0
	for i := 0; i < n; i++ {
		if arr[i]%2 == 0 {
			l = i
			break
		}
	}
	//fmt.Println(l)
	// сумма чисел масива
	sum := 0
	if l < min {
		for i := l; i <= min; i++ {
			sum += arr[i]
		}
	} else {
		for i := min; i <= l; i++ {
			sum += arr[i]
		}
	}
	fmt.Println(sum)
}

/*
package main

import "fmt"

func main() {
	//запись в массив
	var arr [100]int
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	//первое минимальное значение
	min := 0
	for i := 0; i < n; i++ {
		if arr[i] < arr[min] {
			min = i
		}
	}
	fmt.Println("индекс min", min)
	//первое четное значение
	chet := 0
	for a := 0; a < n; a++ {
		if arr[a]%2 == 0 {
			for ; a < n; a++ {
				chet = arr[a]
			}
		}
	}
	fmt.Println("индекс chet", chet)
	// сумма чисел масива
	sum := 0
	if chet < min {
		for i := chet; i <= min; i++ {
			sum += arr[i]
		}
	} else {
		for i := min; i <= chet; i++ {
			sum += arr[i]
		}
	}
	fmt.Println(sum)
}
*/
