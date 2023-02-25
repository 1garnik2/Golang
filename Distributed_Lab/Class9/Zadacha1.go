/*
Задача 1.
Створити структуру, що містить поля:  прізвище, оклад.
За допомогою цієї структури визначити 5 співробітників з окладами 10000,
12000, 14000, 16000, 20000 і довільними прізвищами. Вивести суму окладів
*/
package main

import "fmt"

type Employee struct {
	Lastname string
	Salary   int
}

func main() {
	worker1 := Employee{"Kononenko", 10000}
	worker2 := Employee{"Shevchenko", 12000}
	worker3 := Employee{"Pilipenko", 14000}
	worker4 := Employee{"Demchenko", 16000}
	worker5 := Employee{"Sidorenko", 20000}
	fmt.Println(worker1.Salary + worker2.Salary +
		worker3.Salary +
		worker4.Salary + worker5.Salary)
}
