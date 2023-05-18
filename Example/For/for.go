package main

import "fmt"

func main() {
	arr := []string{"a", "b", "c"}
	for _, v := range arr {
		fmt.Print(v, " ")
	}
}
