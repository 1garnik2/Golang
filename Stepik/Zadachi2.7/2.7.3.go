package main

import "fmt"

func main() {
	var text string

	_, err := fmt.Scan(&text)
	if err != nil {
		panic("invalid input values")
	}
	ru := []rune(text)
	max := ru[0]
	for i := 1; i < len(ru); i++ {
		per := ru[i]
		if per > max {
			max = per
		}
	}
	fmt.Print(string(max))
}
