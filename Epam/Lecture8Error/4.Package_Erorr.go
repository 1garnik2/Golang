package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero = errors.New("divide dy zero")
var ErrBigNumber = errors.New("Veri big integer number")

func Divide(a, b int) (int, error) {
	if a > 2147483647 || b > 2147483647 {
		return 0, ErrBigNumber
	}
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}

func main() {
	a, b := 8, 100000000000000
	result, err := Divide(a, b)
	if err != nil {
		switch {
		case errors.Is(err, ErrDivideByZero):
			fmt.Println("Divide by zero")
		case errors.Is(err, ErrBigNumber):
			fmt.Println("Big number")
		default:
			fmt.Printf("Unckown error: %v\n", err)
		}
	} else {
		fmt.Printf("%d / %d = %d", a, b, result)
	}
}
