package main

import (
	"fmt"
	"math"
)

func main() {
	var n float64
	n = 12345
	a := math.Mod(n, 10)
	b := math.Mod((n-a)/10, 10)
	c := math.Remainder((n-(b*10)-a)/100, 10)
	d := math.Mod((n-((c*100)+(b*10)+a))/1000, 10)
	e := math.Remainder((n-((d*1000)+(c*100)+(b*10)+a))/10000, 10)
	fmt.Print(a*10000 + b*1000 + c*100 + d*10 + e)
}
