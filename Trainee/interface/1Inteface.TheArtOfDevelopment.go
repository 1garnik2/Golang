package main

import "fmt"

type struckHere struct {
	N1, N2 int
}

type interfaceHere interface {
	Sum() int
	Sub() int
}

func (s *struckHere) Sum() int {
	return s.N1 + s.N2 
}

func (s *struckHere) Sub() int {
	return s.N1 - s.N2
}

func main() {
	var a, b interfaceHere //
	sh := struckHere{4, 7}
	os := otherStruct{2, 6}
	a = &sh
	b = os
	fmt.Println(a.Sum())
	fmt.Println(a.Sub())
	fmt.Println(b.Sum())
	fmt.Println(b.Sub())

}

type otherStruct struct {
	A, B int
}

func (o otherStruct) Sub() int {
	return o.A - o.B
}
func (o otherStruct) Sum() int {
	return o.A + o.B
}
