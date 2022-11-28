package main

import "fmt"

func main() {
	var x int = 100
	var p *int

	p = &x
	fmt.Println("Values stored in x", x)
	fmt.Println("Adress of x =", &x)
	fmt.Println("Value stored in variable p =", p)
	fmt.Println("Value stored in x(*p)", *p)
	*p = 200
	fmt.Println("Value stored in variable p =", p)
	fmt.Println("Value stored in x(*p)", *p)
	fmt.Println("value stored in x=", x)
	var p2 *int
	fmt.Println("Value stored in variable p2=", p2)
	var p3 = new(int)
	fmt.Println("value stored in variable p3=", p3)
	fmt.Println("Value stored in *p3", *p3)
	var p4 = &p
	fmt.Println("Value stored in variable p4", p4)
	fmt.Println("Value stored in *p4", *p4)
	fmt.Println("Value stored in **p4", **p4)
}
