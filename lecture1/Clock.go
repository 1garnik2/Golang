/*
Программа высчитывает часы и минуты
в зависимости от поворота часовой стрелки
на определенный граду 0 - 360
*/

package main

import "fmt"

func main() {
	var h, m, t int32

	i := "It is"
	b := "hours"
	d := "minutes."
	fmt.Scanf("%d", &t)
	h = t / 30
	m = (t % 30)
	fmt.Println(i, h, b, m, d)
}
