package main

import "fmt"

/* import "fmt"

func change(str *string) {
	*str = "LOL"
}
func main() {
	s := "LLL"
	fmt.Println(s)
	var poiner *string = &s
	change(poiner)
	fmt.Println(s)
}
*/
func main() {
	num := 9
	b := &num
	*b = 35
	fmt.Println(*b, num)
}
