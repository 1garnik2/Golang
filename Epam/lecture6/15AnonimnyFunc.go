package main

import "fmt"

func main() {

	f := func(x int, y int) int { return x + y }
	fmt.Println(f(4, 8))
	fmt.Println(f(45, 78))

	f = func(x, y int) int { return x - y }
	fmt.Println(f(7, 3))
	action(10, 25, func(x, y int) int { return x + y }) //35

	action(4, 5, func(i1, i2 int) int { return i1 * i2 })

	f = selectFn(1)
	fmt.Println(f(4, 8)) //12

	f = selectFn(2)
	fmt.Println(f(12, 7))

	f = selectFn(3)
	fmt.Println(f(78, 43))

}
func action(i1 int, i2 int, operation func(int, int) int) {
	result := operation(i1, i2)
	fmt.Println(result)
}

func selectFn(n int) func(int, int) int {
	if n == 1 {
		return func(x, y int) int { return x + y }

	} else if n == 2 {
		return func(x, y int) int { return x - y }

	} else {
		return func(x, y int) int { return x * y }
	}
}
