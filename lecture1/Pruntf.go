package main

import "fmt" //пакет fmt

func main() {
	x, y := 2, 4
	s, g := "EPAM", "GO"
	fmt.Printf("x = %d y = %d \n", x, y) // тип данных int
	fmt.Printf("%s   %s \n ", s, g)      // тип данных string
}
