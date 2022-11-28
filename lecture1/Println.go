package main

import "fmt"

func main() {
	x, y := 2, 4
	s, g := "EPAM", "GO"
	fmt.Println(x, y) // Print НЕ ставит пробелы в сторках
	fmt.Println(s, g) // Println сам ставит пробелы в строках

}
