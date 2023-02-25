package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	if b := 10; a > b {
		fmt.Println("Value", a, "more than 10")
	} else {
		fmt.Println("Value", a, "not more than 10")
	}
}
