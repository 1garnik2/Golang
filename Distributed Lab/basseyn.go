package main

import (
	"fmt"
	"math"
)

func main() {
	var n, k, hour1, hour2, bas1, bas2, Zahour, day float64
	fmt.Print("Введіть за скількі днів перша труба наповнить бассейн:")
	fmt.Scan(&n)
	fmt.Print("Введіть за скількі днів друга труба наповнить бассейн:")
	fmt.Scan(&k)
	hour1 = n * 24
	hour2 = k * 24
	bas1 = 1 / hour1
	bas2 = 1 / hour2
	Zahour = bas1 + bas2
	day = (1 / Zahour)
	f := int(math.Round(day))
	d := f / 24
	h := f % 24
	fmt.Printf("Бассейн буде наповнений за %d дня i %d годин", d, h)

}
