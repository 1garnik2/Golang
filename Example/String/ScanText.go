package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// var 1
	var s string

	myscanner := bufio.NewScanner(os.Stdin)
	myscanner.Scan()
	s = myscanner.Text()

	fmt.Println(s)
}
