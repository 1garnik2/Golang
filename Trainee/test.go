package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func main() {
	var t = 45
	fmt.Println(t)

}
func scrabPage(url string) {
	c := colly.NewCollector()

	c.OnHTML()

}
