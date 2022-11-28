//Створити  функцію, яка знаходить кількість
// цифр цілого додатного числа .
//  Використовуючи цю функцію,
// знайти кількість цифр для чисел: 321,
//1243, 12543.

package main

import (
	"fmt"
)

var simvol string

func Kol(chyslo int) int {
	count := 0
	for chyslo > 0 {
		chyslo = chyslo / 10
		count++
	}

	return count
}

func main() {

	a := Kol(321)
	b := Kol(1234)
	c := Kol(12543)
	fmt.Println(a, b, c)

}

/*
Задача 3 версія2
package main

import (
    "fmt"
    "strconv"
)

func Kol(chyslo int) int {
    simvol := strconv.Itoa(chyslo)
    return len(simvol)
}

func main() {
    a := Kol(321)
    b := Kol(1234)
    c := Kol(12543)
    fmt.Println(a, b, c)

}


*/
