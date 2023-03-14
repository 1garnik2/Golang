package main

import "fmt"

func main() {
	year := 1900
	fmt.Println(century(year))
}
func century(year int) int {
	centur := 0
	if year > 0 {
		centur = year / 100
		if year%100 > 0 {
			centur += 1
		}
	}

	return centur
}
