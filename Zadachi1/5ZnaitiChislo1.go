package main

import "fmt"

func main() {
	var i int32
	fmt.Scanf("%d", &i)
	fmt.Println(i%10*100 + i/10)
	/*
		Число 789
		% остаток от деления на 10 равен 9
		+ на цело делим (789/10) получаем 78
		9*100 + 78
		вывод:978
	*/

}
