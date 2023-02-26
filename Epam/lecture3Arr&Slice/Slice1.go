package main

import "fmt"

func main() {
	// инициализация слайса
	var intSlice = make([]int, 10)
	var strSlice = make([]string, 10, 20)

	fmt.Printf("intSlice \tLen: %v \tCap %v\n", len(intSlice), cap(intSlice))
	fmt.Println(strSlice)

	fmt.Printf("strSlice \tLen: %v \tCap: %v\n", len(strSlice), cap(strSlice))
	fmt.Println(strSlice)
	fmt.Printf("%q\n", strSlice)
	// инициализация слайса
	var intSlice2 = new([50]int)[0:10]
	fmt.Printf("intSlice2 \tLen: %v \tCap: %v\n", len(intSlice2), cap(intSlice2))
	fmt.Println(intSlice2)
}
