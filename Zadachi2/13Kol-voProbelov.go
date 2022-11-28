package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var s string
	myscanner := bufio.NewScanner(os.Stdin)
	myscanner.Scan()
	s = myscanner.Text()

	fmt.Println(strings.Count(s, " "))
}
