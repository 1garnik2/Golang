// конвертация в рядок или конвертация из рядка
package main

import (
	"fmt"
	"strconv"
)

func main() {
	//конвертация целого числа в рядок
	s := strconv.Itoa(123)
	fmt.Printf("%s %T\n", s, s)
	//конвертация рядка в целое число
	n, err := strconv.Atoi("123")
	fmt.Println(err)
	fmt.Printf("%d %T\n", n, n)
	//Нельзя конвертировать буквы в цифры ОШИБКА
	n, err = strconv.Atoi("abc")
	fmt.Println(err)
	fmt.Printf("%d %T\n", n, n)
	//форматирование числа в рядок в десятичной системе исчисления
	s = strconv.FormatInt(123, 10)
	fmt.Printf("%s %T\n", s, s)
	//форматирование числа в рядок в двоичной системе исчисления
	s = strconv.FormatInt(123, 2)
	fmt.Printf("%s %T\n", s, s)
	//форматирование рядка в число в десятичной системе исчисления
	//64 это разрядность
	x, err := strconv.ParseInt("123", 10, 64)
	fmt.Printf("%d %T\n", x, x)
	//форматирование рядка в число в двоичной системе исчисления
	//64 это разрядность
	y, err := strconv.ParseInt("1111011", 2, 64)
	fmt.Printf("%d %T\n", y, y)
}
