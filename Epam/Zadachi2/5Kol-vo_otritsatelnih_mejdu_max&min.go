package main

import "fmt"

func main() {
	var arr [100]int
	var n, count, l, r int
	fmt.Scan(n)

	for i := 0; i < n; i++ {
		fmt.Scan(arr[i])
	}
	//Найти позицию первого минимального и
	// последнего максимального
	for i := 1; i < n; i++ {
		if arr[i] < arr[l] {
			l = i
		}
		if arr[i] >= arr[r] {
			r = i
		}
	}
	//
	if l > r {
		l, r = r, l
	}
	for i := l; i <= r; i++ {
		if arr[i] < 0 {
			count++
		}
	}
	fmt.Println(count)

}
