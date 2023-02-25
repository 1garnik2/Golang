package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Я люблю програмування на мові C++"
	text = strings.Replace(text, "C++", "Go", 1)
	fmt.Printf("%s\n", text)
	text = strings.ToUpper(text)
	fmt.Printf("%s \n", text)
	text = strings.Repeat(text, 10)
	fmt.Printf("%s\n", text)
	fmt.Printf("Я-%d Л-%d Ю-%d Б-%d П-%d Р-%d О-%d Г-%d А-%d М-%d У-%d В-%d Н-%d І-%d G-%d O-%d",
		strings.Count(text, "Я"), strings.Count(text, "Л"),
		strings.Count(text, "Ю"), strings.Count(text, "Б"),
		strings.Count(text, "П"), strings.Count(text, "Р"),
		strings.Count(text, "О"), strings.Count(text, "Г"),
		strings.Count(text, "А"), strings.Count(text, "М"),
		strings.Count(text, "У"), strings.Count(text, "В"),
		strings.Count(text, "Н"), strings.Count(text, "І"),
		strings.Count(text, "G"), strings.Count(text, "O"))

}
