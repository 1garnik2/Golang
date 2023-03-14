/*
Напишите функцию sumInt, принимающую переменное количество аргументов типа int, и возвращающую количество полученных функцией аргументов и их сумму. Пакет "fmt" уже импортирован, функция и пакет main объявлены.

Пример вызова вашей функции:

a, b := sumInt(1, 0)
fmt.Println(a, b)

Результат: 2, 1
*/
package main

func main() {
	sumInt(1, 0)

}

func sumInt(a ...int) (int, int) {
	var sum int
	for i := 0; i < len(a); i++ {
		sum += a[i]
	}
	return len(a), sum
}
