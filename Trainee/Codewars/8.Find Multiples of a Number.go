package main

import "fmt"

func main() {
	var a, b = 5, 25          

	fmt.Println(FindMultiples(a, b))

}

func FindMultiples(integer, limit int) []int {
	c := integer
	slice := []int{}
	for i := integer; i <= limit; i = i + c {
		slice = append(slice, i)
	}
	return slice

}
  