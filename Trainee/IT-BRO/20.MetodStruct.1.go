package main

import "fmt"

type User struct {
	name     string
	age      int
	password string
}

//метод который возвращает значение
func (u User) getName() string {
	return u.name
}

//метод который устанавливает новое значение
func (u *User) setName(name1 string) {
	u.name = name1
}
func main() {
	user := User{"Robert", 34, "loginpass"}
	fmt.Println(user.getName())

	user.setName("Barabos")
	fmt.Println(user.name)
}
