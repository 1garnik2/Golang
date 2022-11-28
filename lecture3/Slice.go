package main

import "fmt"

func main() {
	mas := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var slice []int
	slice = mas[1:4] // выводит с 1 по 4 не включительно!

	fmt.Print("Print Slice: ")
	for _, i := range slice {
		fmt.Print(i, " ")
	}
	fmt.Println()

	slice[0] = 6
	fmt.Print("Print array:")
	for _, i := range mas {
		fmt.Print(i, " ")
	}
	fmt.Println()

	mas[2] = 15
	fmt.Print("Print Slice:")
	for _, i := range slice {
		fmt.Print(i, " ")
	}
	fmt.Println()

	slice = append(slice, 16)
	fmt.Print("Print Slice: ")
	for _, i := range slice {
		fmt.Print(i, " ")
	}
	fmt.Println()

	fmt.Print("Print array:")
	for _, i := range mas {
		fmt.Print(i, " ")
	}
	fmt.Println()

	var slice2 []int
	slice2 = slice[1:3]
	fmt.Print("Print Slice2:")
	for _, i := range slice2 {
		fmt.Print(i, " ")
	}
	fmt.Println()

	slice2[1] = 20
	fmt.Println(mas)
	fmt.Println(slice)
	fmt.Println(slice2)
}
