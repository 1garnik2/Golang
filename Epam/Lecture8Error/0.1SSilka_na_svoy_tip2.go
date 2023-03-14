// Линейный двухсвязный список (структура данных)
package main

import "fmt"

type emploee struct {
	name        string
	age         int
	designation string
	salary      int
	chief       *emploee
	pidlegiy    *emploee
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
	emploee2.pidlegiy = &emploee1
	fmt.Println(emploee1.chief.name)

	emploee3 := emploee{
		name:        "Robert Griesemer",
		designation: "Director",
	}
	emploee2.chief = &emploee3
	emploee3.pidlegiy = &emploee2
	fmt.Println(emploee1.chief.chief.name)
	fmt.Println(emploee2.chief.name)
	fmt.Println(emploee3.pidlegiy.pidlegiy.name)
}
