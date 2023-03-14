//

package main

import "fmt"

func main() {
	arr := [4]int{6, 2, 7, 3}
	s := arr[:]
	fmt.Println(decode(s, 4))
}
func decode(encoded []int, first int) []int {
	XOR := make([]int, len(encoded)+1)
	XOR[0] = first
	for i := 0; i < len(XOR)-1; i++ {

		XOR[i+1] = XOR[i] ^ encoded[i]

	}
	return XOR
}
