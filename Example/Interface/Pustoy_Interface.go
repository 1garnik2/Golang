package main

import (
	"fmt"
	"reflect"
)

func main() {
	m := map[string]interface{}{}
	m["One"] = 1
	m["two"] = 2.0
	m["three"] = true

	for k, v := range m {
		switch v.(type) {
		case int:
			fmt.Printf("%s is an integer \n", k)
		case float64:
			fmt.Printf("%s is an float64 \n", k)
		default:
			fmt.Printf("%s is %v", k, reflect.TypeOf(v))
		}
	}
}
