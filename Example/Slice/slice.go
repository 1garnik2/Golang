package main

import "fmt"

func main() {
	//создание слайсов разных типов
	var intSlice = make([]int, 10)
	var strSlice = make([]string, 10, 20)

	//вывод слайса
	fmt.Printf("intSlice \t len: %v \tCap:%v\n", len(intSlice), cap(intSlice))
	fmt.Println(intSlice)
	//вывод слайса
	fmt.Printf("strSlice \t len: %v \tCap:%v\n", len(strSlice), cap(strSlice))
	fmt.Println(intSlice)
	fmt.Printf("%q\n", strSlice)
	//инициализация нового слайса
	var intSlice2 = new([50]int)[0:10]
	fmt.Printf("intSlice2 \tlen: %v \t Cap: %v\n", len(intSlice2), cap(intSlice2))
	fmt.Println(intSlice2)

}
