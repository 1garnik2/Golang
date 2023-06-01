package main

import "fmt"

type Dat struct {
	age int
	kol int
}

func main() {
	money := map[string]Dat{
		"dollar": {3000, 67},
		"euro":   {4444, 34},
		"apples": {90, 33},
	}
	for key, val := range money {
		fmt.Println(val.age)
		fmt.Println(key)
	}
	//del
	delete(money, "apples")
	//Add
	money["kivi"] = Dat{68, 798}
	//Update
	dollar := money["dollar"]
	dollar.kol = 99
	money["dollar"] = dollar
	fmt.Println(money)

	status, ok := money["dollar"]
	if ok {
		fmt.Println("bool")
	}
	fmt.Println(status, ok)

}
