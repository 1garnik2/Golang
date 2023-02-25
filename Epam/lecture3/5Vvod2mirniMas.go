package main

import "fmt"

func main() {
	var mas [3][3]int

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Scan(&mas[i][j])
		}
	}
	for i := range mas {
		fmt.Println(mas[i])
	}
	sum := 0

	for _, row := range mas {
		for _, x := range row {
			sum += x
		}
	}
	fmt.Println("sum =", sum)

	sumd := 0
	for i := 0; i < 3; i++ {
		sumd += mas[i][i]
	}
	fmt.Println("sum diagonal =", sumd)

	sumpd := 0
	for i := 0; i < 3; i++ {
		sumpd += mas[i][2-i]
	}
	fmt.Println("Sum Pobochnoy diagonaly =", sumpd)
}
