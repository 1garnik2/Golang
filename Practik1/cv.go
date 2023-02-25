// считает количество нулей в строке, цифре или ищет одинаковые символы
package main

import (
	"fmt"
	"strings"
)

func main() {
	var chislo string
	fmt.Scan(&chislo)
	fmt.Println(strings.Count(chislo, "0"))
}
