package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a) // сканирование и запись переменной
	fmt.Scan(&b)
	if a > b {
		fmt.Println(a)
	} else {
		fmt.Println(b)
	}
}
