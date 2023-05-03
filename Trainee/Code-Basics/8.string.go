package main

import (
	"fmt"
	"strings"
)

func Greetings(name string) string {
	name = strings.Trim(name, " ")
	name = strings.ToLower(name)
	name = strings.Title(name)
	return fmt.Sprintf("Привет, %s!", name)
}

// BEGIN (write your solution here)
func main() {
	var a string
	fmt.Scan(&a)
	fmt.Println(Greetings(a))
}

// END
