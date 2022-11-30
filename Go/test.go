package main

import (
	"fmt"
)

func main() {
	emoji := []rune("привет")

	for i := 0; i < len(emoji); i++ {
		fmt.Print(string(emoji[i]), " ") // выводим код символа и его строковое представление
	}
}
