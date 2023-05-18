package main

import "fmt"

type myInt int
type myStr string

func (i myInt) String() string {
	return fmt.Sprintf("%d", i)
}
func (s myStr) String() string {
	return string(s)
}

//                           (передаваемый
//            (ограничения)    слайс или.) (возврат)
func ToStrings[T fmt.Stringer](elems []T) []string {
	r := make([]string, len(elems))

	for i, v := range elems {
		r[i] = v.String()
	}
	return r
}

func main() {
	arr := [...]myInt{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arr2 := [...]myStr{"a", "d", "g", "f", "r", "d", "t", "h", "k", "f", "x", "f"}
	str2 := ToStrings(arr2[:])
	for _, v := range str2 {
		fmt.Printf("Value str2 %s\n", v)
	}
	fmt.Println(str2)
	str := ToStrings(arr[:])
	for i, v := range str {
		fmt.Printf("value:%s type:%T index:%d\n", v, v, i)
	}
}
