package main

import "fmt"

func main() {
	var friendOfDima []string
	friendsOfSemyon := make([]string, 3)
	friendOfDima = append(friendOfDima, "Костя", "Семен", "Таня")
	friendsOfSemyon = append(friendsOfSemyon, "Валера", "Таня", "Дима")
	friends := map[string][]string{
		"Dima":   friendOfDima,
		"Semyon": friendsOfSemyon,
		"Костя":  nil,
	}
	_, value := friends["Костя"]
	fmt.Print(value, " ")
	delete(friends, "Dima")
	fmt.Print(friendsOfSemyon[3])
}
