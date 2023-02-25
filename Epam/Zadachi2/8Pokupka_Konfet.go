// полный перебор
package main

import "fmt"

func main() {

	var arr [100]int
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	max := 0
	A := 0
	B := 0
	// for i := 0; i < n-3; i++ перебор левой рукой
	for i := 0; i < n-3; i++ {
		// перебор правой рукой
		for j := i + 2; j < n-1; j++ {
			if arr[i]+arr[i+1]+arr[j]+arr[j+1] > max {
				max = arr[i] + arr[i+1] + arr[j] + arr[j+1]
				A = i + 1
				B = j + 1

			}
		}

	}
	fmt.Println(A, A+1, B, B+1)
}
