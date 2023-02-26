package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "34"
	d, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(d)
	}

}
