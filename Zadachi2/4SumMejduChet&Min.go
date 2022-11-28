package main

import "fmt"

func main() {
	var arr [100]int
	var n, sum, l, r int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	for i := 0; i < n; i++ {
		if arr[i]%2 == 0 {
			l = i
			break
		}
	}
	for i := 0; i < n; i++ {
		if arr[i] < arr[r] {
			r = i
		}
	}
	//перезапись єлементов
	if l > r {
		l, r = r, l
	}
	for i := l; i < r; i++ {
		sum += arr[i]
	}
	fmt.Println(sum)
}
