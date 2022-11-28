package main // складывает квадраты наибольшего и наименьшего

import (
	"fmt"
)

func main() {
	var a, b, c, m, q int32

	fmt.Scan(&a, &b, &c)
	if a >= b && a >= c {
		m = a
	} else {
		if b >= a && b >= c {
			m = b
		} else {
			m = c
		}
	}
	if a <= b && a <= c {
		q = a
	} else {
		if b <= a && b <= c {
			q = b
		} else {
			q = c
		}
	}
	fmt.Println(m * m)
	fmt.Println(q * q)
}
