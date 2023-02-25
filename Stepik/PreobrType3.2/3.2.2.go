package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var a, b float64
	c, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	c = strings.Replace(c, " ", "", -1)
	c = strings.Replace(c, ",", ".", -1)
	fmt.Sscanf(c, "%f;%f", &a, &b)
	fmt.Printf("%.4f\n", a/b)
}
