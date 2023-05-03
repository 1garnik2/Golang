package main

import "fmt"

type User struct {
	name     string
	age      int
	password string
	score    []int
}

func (u User) getHighScore() int {
	high := 0
	for _, v := range u.score {
		if high < v {
			high = v
		}
	}
	return high
}

func main() {
	user := User{"Bob", 22, "Pas", []int{23, 67, 1902, 345, 1}}
	fmt.Println(user.getHighScore())
	fmt.Println(user)
}
