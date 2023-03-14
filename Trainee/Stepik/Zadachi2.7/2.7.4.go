package main

import (
	"fmt"
	"strconv"
)

func main() {
	var chislo string
	fmt.Scan(&chislo)

	for _, r := range chislo {
		did, err := strconv.Atoi(string(r))
		if err != nil {
			panic(fmt.Sprintf("Error %c", r))
		}
		fmt.Print(did * did)
	}
}
