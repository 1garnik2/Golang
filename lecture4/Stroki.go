package main

func main() {
	var str1 = "Hello, Epam!!"
	println(str1)

	var str2 = "привіт , Epam!"
	println(str2)

	println(str1[0])         // код таблицы аски
	println(string(str1[0])) // вывод буквы

	for i := 0; i < 5; i++ {
		print(string(str1[i]))
	}
	println()

	println(str2[0])

	// с кирилицей не работае!
	for i := 0; i < 6; i++ {
		print(string(str2[i]))

	}
	println()
	// с кирицей работает!!!
	for _, i := range str2 {
		print(string(i))
	}
}
