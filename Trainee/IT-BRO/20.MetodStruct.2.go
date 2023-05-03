package main

import "fmt"

type Party struct {
	name     string
	age      int
	password string
}

func (u Party) isElder() bool {
	if u.age < 18 {
		return false
	}
	return true

}

func main() {
	user := Party{"Kate", 17, "pfss"}
	if user.isElder() {
		fmt.Println("Заходи!")
	} else {
		fmt.Println("Нельзя!")
	}
	fmt.Println(user)
}
