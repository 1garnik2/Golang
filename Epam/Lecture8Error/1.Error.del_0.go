package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	c, err := divide(x, y)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(c)
	}

}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("can't didvide '%d' by zero", a)
	}
	return a / b, nil
}
