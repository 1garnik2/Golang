package main

import "fmt"

func main() {
	var ch = 1200000
	j := 0
	for i := 1; i < ch; i++ {
		if i%2 == 0 || i%5 == 0 && (i%2 != 0 || i%5 != 0) {
			j = j + 1
		}

	}
	one := 1200000 / 100
	ost := j / one

	fmt.Println(ost, "%")
}
