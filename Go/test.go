package main

import "fmt"

type Struct struct {
	On          bool
	Ammo, Power int
}

func (obj *Struct) Shoot() bool {
	if obj.On == false || obj.Ammo == 0 {
		return false
	}
	obj.Ammo--
	return true
}

func (obj *Struct) RideBike() bool {
	if obj.On == false || obj.Power == 0 {
		return false
	}
	obj.Power--
	return true
}

func main() {
	testStruct := Struct{true, 2, 0}
	fmt.Println(&testStruct)
}
