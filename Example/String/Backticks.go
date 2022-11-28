package main

import "fmt"

func main() {
	var str1 = "Hello, \n Epam"
	fmt.Println(str1)
	//конкотенация
	var str = "Hello," +
		" Epam"
	fmt.Println(str)
	// или
	var str2 = `Hello,Epam` //backticks``
	fmt.Println(str2)

	var str3 = `Hello,
	 Epam` //backticks``
	fmt.Println(str3)

	var str4 = `Hello, \n Epam` //backticks``
	fmt.Println(str4)

}
