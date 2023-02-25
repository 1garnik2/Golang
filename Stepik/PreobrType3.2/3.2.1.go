package main

import "fmt"

func main() {
	var i int64
	i = 5765
	fmt.Print(convert(i))

}
func convert(v int64) uint16 {
	return uint16(v)
}
