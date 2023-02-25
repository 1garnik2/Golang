package main

import "fmt"

type contact struct {
	email string
	phone string
}

type rabotnik struct {
	name        string
	age         int
	designation string
	salary      int
	contact
}

func main() {
	rabotnik1 := rabotnik{
		name:        "Ken Thompson",
		age:         44,
		designation: "Tester",
		salary:      1800,
		contact: contact{
			email: "123@gh.mail.com",
			phone: "123-123-234",
		},
	}
	fmt.Println(rabotnik1)
	fmt.Println(rabotnik1.contact)
	fmt.Println(rabotnik1.contact.email)
	//можно указывать без промежуточной структуры
	fmt.Println()
	fmt.Println(rabotnik1.email)
	fmt.Println(rabotnik1.phone)
}
