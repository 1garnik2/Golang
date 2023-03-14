package main

import "fmt"

func main() {
	slice := []int{1, 2, 2}
	fmt.Println(SquareSum(slice))
}

func SquareSum(numbers []int) int {
	// your code here
	num := 0
	for _, v := range numbers {
		num += v * v
	}
	return num
}
