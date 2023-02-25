// PowerOfTwo is binary shift
package main

import "fmt"

func main() {
	var a = 7
	var b = 5
	b = a ^ b
	//var c int
	a = a >> b
	fmt.Println(a)
	fmt.Println(b)

	//a=0111
	//b=0101
	//a^b =0011
}
func PowerOfTwo(v int) int {
	v--
	v |= v >> 1
	v |= v >> 2
	v |= v >> 4
	v |= v >> 8
	v |= v >> 16
	v++

	return v
}
