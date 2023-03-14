/*Напишите функцию, которая умножает
значения на которые ссылаются два указателя
и выводит получившееся произведение в консоль.
 Ввод данных уже написан за вас.

Sample Input:

2 2
Sample Output:

4
Напи*/package main

import "fmt"

func main() {
	a := 11
	b := 2
	test(&a, &b)

}
func test(x1 *int, x2 *int) {
	fmt.Println(*x1 * *x2)

}
