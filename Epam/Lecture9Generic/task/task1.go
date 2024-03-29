package main

import "fmt"

func fibonacciLoop[T ~int](n T) []T {
	f := make([]T, n+1, n+2)
	if n < 2 {
		f = f[0:2]
	}
	f[0] = 0
	f[1] = 1
	for i := 2; T(i) <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f
}
func main() {
	fmt.Println(fibonacciLoop(10))
}
