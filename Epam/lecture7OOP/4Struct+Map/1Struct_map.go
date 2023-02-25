package main

import "fmt"

type employee struct {
	name        string
	age         int
	designation string
	salary      int
}
type deportament struct {
	name string
	room int
}

func main() {
	var Rob = employee{"Rob Pyke", 33, "Developer", 2000}
	var Ken = employee{"Ken Thompson", 23, "Tester", 1800}
	var Robert = employee{"Robert Griesemer", 30, "TeamLead", 4800}

	var dep1 = deportament{"Golang", 210}
	var dep2 = deportament{"Java", 215}

	var mapa = map[employee]deportament{}
	//запись ключа(имя) и значения(депортамент)
	mapa[Rob] = dep1
	mapa[Ken] = dep2
	mapa[Robert] = dep1
	//

	/////////////////////////ключ..значение//////
	fmt.Println(Ken.age, mapa[Rob].room)
	fmt.Println(mapa)              // Вывод карты
	fmt.Println(mapa[Robert].name) //Вывод названия отдела Rob
	fmt.Println(mapa[Rob].room)    //Вывод номера комнаты Rob
	//==========================================
	// при измении значения в стрктуре Rob
	// ключ для мапы измениться
	// и при выводе мапы будет 0 как в примере ниже
	Rob.age = 22
	fmt.Println(mapa[Rob].room)

}
