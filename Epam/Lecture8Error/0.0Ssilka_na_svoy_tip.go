// Линейный односвязный список (структура данных)
package main

import "fmt"

type emploee struct {
	name        string
	age         int
	designation string
	salary      int
	chief       *emploee
}

func main() {
	emploee1 := emploee{
		name:        "Rob Pyke",
		age:         25,
		designation: "Developer",
		salary:      2000,
	}
	fmt.Println(emploee1)

	emploee2 := emploee{
		name:        "Ken Thompson",
		age:         30,
		designation: "Manager",
		salary:      4000,
	}
	emploee1.chief = &emploee2
	fmt.Println(emploee1.chief.name)

	emploee3 := emploee{
		name:        "Robert Griesemer",
		designation: "Director",
	}
	emploee2.chief = &emploee3
	fmt.Println(emploee1.chief.chief.name)
	fmt.Println(emploee2.chief.name)
}
