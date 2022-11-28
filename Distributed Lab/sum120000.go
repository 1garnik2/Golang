package main

import "fmt"

func main() {

	j := 1200000
	sum := 0
	for i := 1; i < j; i++ {
		if (i%2 == 0 || i%3 == 0) && (i%2 != 0 || i%3 != 0) {

			sum = sum + i
		}
	}
	fmt.Println(sum)
}
