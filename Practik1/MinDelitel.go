package main

import "fmt"

func main() {
	var i, buf int
	fmt.Scan(&i)
	a := (i % 100)
	for ; i > 1; i-- {
		if a%i == 0 {
			buf = i

		}

	}
	fmt.Println(buf)
}
