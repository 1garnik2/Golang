// Больше соседних
package main

import "fmt"

func main() {
	//заполнение массива на 10 элементов
	var arr [100]int
	var n, count int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	for i := 1; i < n; i++ {
		if arr[i] > arr[i-1] && arr[i] > arr[i+1] {
			count++
		}
	}
	fmt.Println(count)
}

/*
package main

import "fmt"

func main() {
	var arr [100]int
	var n int
	fmt.Scan(&n)
	for i := 0; i < 100; i++ {
		fmt.Scan(&arr[i])
	}
	temp := 0
	for i := 1; i < (n - 1); i++ {
		if arr[i] > arr[i-1] && arr[i] > arr[i+1] {
			temp += 1
		}

	}

	fmt.Print(temp)

}*/
