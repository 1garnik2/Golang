package main

import "fmt"

func main() {

	var str1 = "bravo"
	fmt.Println(str1, len(str1))

	var str2 = "браво"
	fmt.Println(str2, len(str2))
	//вывод номера в массиве,код АСКИ, буквы
	fmt.Println("String 1")
	for i, c := range str1 {
		fmt.Println(i, c, string(c))
	}
	//вывод номера в массиве,код АСКИ, буквы
	//начинается с парного байта 0 2 4 6 8
	fmt.Println("Stribg 2")
	for i, c := range str2 {
		fmt.Println(i, c, string(c))
	}
	fmt.Println()
}
