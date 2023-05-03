// создание структуры и изменение поля name через указатели
package main

import "fmt"

type User struct {
	name     string
	age      int
	password string
}

func change(u *User) {
	u.name = "Kate"
}

func main() {
	user := User{"Bob", 43, "pass"}
	fmt.Println(user)
	change(&user)
	fmt.Println(user)
}
