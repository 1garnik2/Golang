package main

import "fmt"

func main() {
	var time = 19
	buc := 1
	for i := 0; i < time; i++ {
		buc = buc * 2
	}
	fmt.Println(buc)
}
