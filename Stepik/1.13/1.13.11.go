/*
По данному числу n закончите фразу "На лугу пасется..."

	одним из возможных продолжений: "n коров", "n корова",
	 "n коровы", правильно склоняя слово "корова".

# Входные данные

Дано число n (0<n<100).

# Выходные данные

Программа должна вывести введенное число n и одно из

	слов (на латинице): korov, korova или korovy,
	например, 1 korova, 2 korovy, 5 korov. Между числом
	и словом должен стоять ровно один пробел.

Sample Input:

10
Sample Output:

10 korov
*/
package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)

	if a == 1 || (a > 20 && a%10 == 1) {
		fmt.Printf("%d korova", a)
	} else if a%10 >= 2 && a%10 <= 4 && (a < 4 || a > 21) {
		fmt.Printf("%d korovy", a)
	} else {
		fmt.Printf("%d korov", a)
	}
}
