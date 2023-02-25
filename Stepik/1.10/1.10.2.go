/*
Требуется написать программу, при выполнении которой
с клавиатуры считываются два натуральных числа A и B

	(каждое не более 100, A < B).
	 Вывести сумму всех чисел от A до B  включительно.

Sample Input:

1 5
Sample Output:

15
*/
package main

import "fmt"

func main() {

	var i, a, b int
	fmt.Scan(&a, &b)

	for ; a <= b; a++ {
		i += a

	}
	fmt.Println(i)
}
