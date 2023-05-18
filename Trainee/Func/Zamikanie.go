package main

import "fmt"

func totalPrice(initPrice int) func(int) int {
	sum := initPrice
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	orderPrice := totalPrice(0)
	for i := 0; i < 5; i++ {
		fmt.Println(orderPrice(1))
	}

}
