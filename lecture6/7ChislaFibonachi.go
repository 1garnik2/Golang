package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Print(fibonacci(i), " ")
	}
	fmt.Println()

}

func fibonacci(number int) int {
	if number == 0 {
		return 0
	}
	if number == 1 {
		return 1
	}
	return fibonacci(number-1) + fibonacci(number-2)
}
