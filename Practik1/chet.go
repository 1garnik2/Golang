package main

import "fmt"

func main() {
	var i int
	fmt.Scan(&i)
	for a := 1; a <= i; a++ {
		for ; a%2 == 0; a++ {
			fmt.Printf("%d ", a)
		}
	}
}
