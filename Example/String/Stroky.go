package main

import "fmt"

func main() {
	var str1 = "Hello, Epam!!"
	fmt.Println(str1)

	var str2 = "привіт , Epam!"
	fmt.Println(str2)

	fmt.Println(str1[0])         //вывод кода из таблицы АСКИ
	fmt.Println(string(str1[0])) //вывод символа Н
	//вывод слова Hello
	for i := 0; i < 5; i++ {
		fmt.Print(string(str1[i]))
	}
	fmt.Println()
	// с кирилицей работает!!!
	for _, i := range str2 {
		print(string(i))
	}

	//такая конструкция Укр язык не выведет!
	fmt.Println(str2[0])
	fmt.Println(string(str2[0]))

}
