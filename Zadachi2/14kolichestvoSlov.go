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
	fmt.Print(len(strings.Fields(s)))

}

/*
package main

import(
	"bufio"
	"fmt"
	"os"
)
func main(){
	var s string
	myscanner:=bufio.NewScanner(os.Stdin)
	myscanner.Scan()
	s=myscanner.Text()
	var count =0
	for i:=0;i<len(s)-1;i++{
		if s[i]!=32 &&s[i+1]==32{
			count++
		}
	}
	fmt.Println(count +1)
}
*/
