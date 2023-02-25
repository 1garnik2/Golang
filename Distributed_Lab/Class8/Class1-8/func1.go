// Створити функцію SumRange(A, B) цілого типу,
// яка знаходить суму всіх цілих чисел від A до B включно
// (A і B — цілі). Якщо A > B, функція повертає 0.
// За допомогою цієї функції знайти суму цілих чисел
// від 10 до 20 (10 і 20 включно).
package main

import "fmt"

func SumRange(A int, B int) (result int) {
	result = (A + B) * (B - A + 1) / 2
	return result
}

func main() {
	value := SumRange(21, 20)
	fmt.Println(value)
}
