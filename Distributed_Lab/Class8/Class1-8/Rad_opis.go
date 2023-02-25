package main

import "fmt"

func main() {
	var chislo int
	fmt.Scan(&chislo)
	if chislo < 10 {
		fmt.Println("Натуральне число")
	}
	if chislo >= 10 && chislo < 100 {
		if chislo%2 == 0 {
			fmt.Println("Парне двозначне число")
		} else {
			fmt.Println("непарне двозначне число")
		}
	}
	if chislo >= 100 && chislo < 1000 {
		if chislo%2 == 0 {
			fmt.Println("парне тризначне число")
		} else {
			fmt.Println("непарне тризначне число")
		}
	}
}
