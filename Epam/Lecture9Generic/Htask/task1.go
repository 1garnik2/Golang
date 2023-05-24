// Task: create chaining methods Add(), Update(), Pop(), Delete(i)
// Cоздайте методы  Add(), Update(), Pop(), Delete(i)      для структуры Struct
package main

import "fmt"

type Struc[T any] struct {
	slice []T
}

func (s *Struc[T]) Add(element T) *Struc[T] {
	s.slice = append(s.slice, element)
	return s
}

// func (s *Struc[T]) Add(element T) {
// 	s.slice = append(s.slice, element)
// }

func (s *Struc[T]) Update(index int, elemet T) *Struc[T] {
	if index >= 0 && index < len(s.slice) {
		s.slice[index] = elemet
	}
	return s
}

func (s *Struc[T]) Pop() *Struc[T] {
	if len(s.slice) > 0 {
		s.slice = s.slice[:len(s.slice)-1]
	}
	return s
}

func (s *Struc[T]) Delete(index int) *Struc[T] {
	if index >= 0 && index < len(s.slice) {
		fmt.Print("del", s.slice[:index])   //все до индекса, индекс не попадает
		fmt.Println(" ", s.slice[index+1:]) // все после индекса,индекс не попадает
		s.slice = append(s.slice[:index], s.slice[index+1:]...)
	}
	return s
}

func main() {
	List := &Struc[int]{}
	List.Add(1).Add(4).Add(8).Add(6).Add(0).Add(7)
	List.Add(2)
	List.Add(7)
	List.Add(8)
	fmt.Println("Add", List.slice)

	List.Update(3, 9).Update(0, 7)
	fmt.Println("Up", List.slice)

	List.Pop()
	fmt.Println("Pop", List.slice)

	List.Delete(3)
	fmt.Println(List.slice)

}
