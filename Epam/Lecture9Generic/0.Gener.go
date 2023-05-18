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

func main() {
	sliceStr := []myStr{"a", "b", "c", "d", "e", "f"}
	sliceInt := []myInt{9, 8, 7, 6, 5, 4, 3, 2, 1}
	for v := range ToStrings(sliceInt) {
		fmt.Print(v, " ")
	}
	fmt.Print("\n")
	str2 := ToStrings(sliceStr)
	for k := range str2 {
		fmt.Print(str2[k], " ")
	}

}

func ToStrings[T fmt.Stringer](elems []T) []string {
	r := make([]string, len(elems))

	for i, v := range elems {
		r[i] = v.String()
	}
	return r
}
