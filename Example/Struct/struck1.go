package main

import "fmt"

type User struct {
	name     string
	age      int
	password string
}

func (u User) getName() string {
	return u.name
}
func main() {
	user := User{"John", 17, "pass"}
	user.name = "Bob"
	fmt.Println(user.password)
	fmt.Println(user.getName())
}
