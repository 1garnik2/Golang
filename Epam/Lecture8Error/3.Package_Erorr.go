package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero = errors.New("Divide by _ zero")

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}
func main() {
	a, b := 6, 0

	res, err := Divide(a, b)
	if err != nil {
		switch {
		case errors.Is(err, ErrDivideByZero):
			fmt.Println("Divide by zero")
		default:
			fmt.Printf("Uncknown error: %v\n", err)
		}
	} else {
		fmt.Printf("%d / %d = %d\n", a, b, res)
	}
}
