/*
складывает самое большое число из 3 цифр
*/
package main

import "fmt"

func main() {
	var a, b, c, d, i, f, g, h, j, k int32
	fmt.Scan(&a)
	b = (a % 100) / 10
	c = a % 10
	d = a / 100

	i = ((b * 100) + (c * 10) + d)
	f = ((c * 100) + (b * 10) + d)
	g = ((d * 100) + (b * 10) + c)
	h = ((b * 100) + (d * 10) + c)
	j = ((c * 100) + (d * 10) + b)
	k = ((d * 100) + (c * 10) + b)
	if i >= f && i >= g && i >= h && i >= j && i >= k {
		fmt.Println(i)
	} else if f >= i && f >= g && f >= h && f >= j && f >= k {
		fmt.Println(f)
	} else if g >= i && g >= f && g >= h && g >= j && g >= k {
		fmt.Println(g)
	} else if h >= i && h >= f && h >= g && h >= j && h >= k {
		fmt.Println(h)
	} else if j >= i && j >= f && j >= g && j >= h && j >= k {
		fmt.Println(j)
	} else if k >= i && k >= f && k >= g && k >= j && k >= h {
		fmt.Println(k)
	}
}
