/*
По данному числу N распечатайте все целые

	значения степени двойки, не превосходящие N,
	в порядке возрастания.

# Входные данные

Вводится натуральное число.

# Выходные данные

Выведите ответ на задачу.

Sample Input:

50
Sample Output:

1 2 4 8 16 32
*/
package main

import "fmt"

func main() {
	var n, s int
	fmt.Scan(&n)
	s = 1
	for s <= n {
		if s > n {
			break
		}
		fmt.Print(s, " ")
		s = s * 2
	}
}
