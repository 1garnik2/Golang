//	Напишите программу для учета оценок студентов в классе.
// Создайте структуру "Студент" (Student)
// с полями "имя" (Name) и "оценки" (Grades),
// где "оценки" - это массив целых чисел.

package main

import "fmt"

type Student struct {
	Name   string
	Grades []int
}

func overGrades(grades []int) float64 {
	sum := 0
	for _, v := range grades {
		sum += v
	}
	return float64(sum) / float64(len(grades))
}

func PrintBall(student []Student) {
	for _, v := range student {
		avg := overGrades(v.Grades)
		fmt.Printf("%s имеет стредний бал %.1f\n", v.Name, avg)
	}
}

func main() {
	var student = []Student{
		{"Alisa", []int{56, 67, 89, 90, 76}},
		{"Vovan", []int{45, 56, 76, 33}},
		{"Daria", []int{77, 45, 98, 65, 45, 67, 89}},
	}
	PrintBall(student)
}
