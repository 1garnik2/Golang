// Выводит оценки школьниов и названия универа и класса
package main

import "fmt"

type Stydenter interface {
	GetMarks() []int
	getUniver() string
}

type Popiler interface {
	GetMarks() []int
	getClass() int
}

//Методы дла получения названия универа или номера класса
func (u Studentstr) getUniver() string {
	return u.univer
}
func (c Popil) getClass() int {
	return c.class
}

// ///////////////////////
type Studentstr struct {
	marks  []int
	univer string
}
type Popil struct {
	marks []int
	class int
}
//Методы дла получения оценок
func (s Studentstr) GetMarks() []int {
	return s.marks
}
func (p Popil) GetMarks() []int {
	return p.marks
}

func main() {
	var stud = Studentstr{[]int{90, 45, 67, 89, 56}, "HGAGH"}
	var Shcool = Popil{[]int{6, 7, 6, 8, 5, 9}, 8}

	var x Stydenter
	x = stud
	fmt.Println(x.GetMarks())

	fmt.Println("Shcool", Shcool.GetMarks())

	fmt.Println(stud.getUniver())
	fmt.Println(Shcool.getClass())
}
