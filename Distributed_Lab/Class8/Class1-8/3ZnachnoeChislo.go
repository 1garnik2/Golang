package main

import "fmt"

func main() {
	var chislo int
	fmt.Scan(&chislo)
	last := chislo / 100
	first := (chislo % 10) * 100
	middle := ((chislo / 10) % 10) * 10
	fmt.Println(first + middle + last)
}
