/*
Внутри функции main (объявлять функцию не нужно) необходимо
написать программу:

На первом этапе на стандартный ввод подается 10 целых
положительных чисел, которые должны быть записаны в порядке
ввода в массив из 10 элементов. Тип чисел,

	входящих в массив, должен соответствовать минимально
	возможному целому беззнаковому числу. Имя массива который
	 вы должны сами создать workArray (условие обязательное).
	  Для чтения из стандартного ввода уже импортирован пакет fmt.

На втором этапе на стандартный ввод подаются еще 3 пары
чисел - индексы элементов этого массива, которые требуется

	 поменять местами (если такая пара чисел 3 и 7, значит
		 в массиве элемент с 3 индексом нужно поменять местами с
		  элементом, индекс которого 7).

Элементы полученного массива должны быть выведены через

	пробел на стандартный вывод. Далее автоматически
	 будет проведена проверка используемых типов, результат
	  которой будет добавлен к вашему ответу.

Использование массива - обязательное условие!

Sample Input:

99 151 137 71 117 187 20 93 187 67 1 2 3 5 7 8
Sample Output:

99 137 151 187 117 71 20 187 93 67 type ok
*/
package main

import "fmt"

func main() {
	// вводиться только содержимое функции!!!!
	var workArray [10]uint8
	var in, in2 uint8

	for i, _ := range workArray {
		fmt.Scan(&workArray[i])
	}

	for i := 0; i < 3; i++ {
		fmt.Scan(&in, &in2)
		workArray[in], workArray[in2] = workArray[in2], workArray[in]
	}
	for _, i := range workArray {
		fmt.Print(i, " ")
	}
}

// abo tac
/*
    var workArray [10]uint8
	var work [6]uint8
	for i := 0; i < 10; i++ {
		fmt.Scan(&workArray[i])
	}
	for i := 0; i < 6; i++ {
		fmt.Scan(&work[i])
	}
	//fmt.Println(work)
	a := work[0]
	b := work[1]
	c := work[2]
	d := work[3]
	e := work[4]
	f := work[5]

	workArray[a], workArray[b] = workArray[b], workArray[a]
	workArray[c], workArray[d] = workArray[d], workArray[c]
	workArray[e], workArray[f] = workArray[f], workArray[e]
	for _, row := range workArray {
		fmt.Print(row, " ")
	}
*/
