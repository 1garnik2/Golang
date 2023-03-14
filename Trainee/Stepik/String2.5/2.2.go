package main

import "fmt"

func main() {
	var s, r string
	fmt.Scan(&s)
	for _, i := range s {
		r = string(i) + r
		fmt.Println(r)
	}
	if s == r {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}
}
