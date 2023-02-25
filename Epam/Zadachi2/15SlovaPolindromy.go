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
	var count = 0
	strSlice := strings.Fields(s)
	var f bool
	for i := 0; i < len(strSlice); i++ {
		f = true
		for j := 0; j < len(strSlice[i])/2; j++ {
			if strSlice[i][j] != strSlice[i][len(strSlice[i])-1-j] {
				f = false
				break
			}
		}
		if f {
			count++
		}
	}
	fmt.Println(count)
}
