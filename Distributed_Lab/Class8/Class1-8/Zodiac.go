package main

import "fmt"

func main() {
	var d, m int
	fmt.Scan(&d, &m)
	if m == 1 && d > 19 || m == 2 && d < 19 { //«Водолій» (20.1–18.2)
		fmt.Println("Водолій")
	}
	if m == 2 && d > 18 || m == 3 && d < 21 { //«Риби» (19.2–20.3),
		fmt.Println("Риби")
	}
	if m == 3 && d > 20 || m == 4 && d < 20 { //«Овен» (21.3–19.4),
		fmt.Println("Овен")
	}
	if m == 4 && d > 19 || m == 5 && d < 21 { // «Телець» (20.4–20.5),
		fmt.Println("Телець")
	}
	if m == 5 && d > 20 || m == 6 && d < 22 { //«Близнюки» (21.5–21.6),
		fmt.Println("Близнюки")
	}
	if m == 6 && d > 21 || m == 7 && d < 23 { //«Рак» (22.6–22.7)
		fmt.Println("Рак")
	}
	if m == 7 && d > 22 || m == 8 && d < 23 { //«Лев» (23.7–22.8)
		fmt.Println("Лев")
	}
	if m == 8 && d > 22 || m == 9 && d < 23 { // «Діва» (23.8–22.9)
		fmt.Println("Діва")
	}
	if m == 9 && d > 22 || m == 10 && d < 23 { //«Терези» (23.9–22.10),
		fmt.Println("Терези")
	}
	if m == 10 && d > 22 || m == 11 && d < 23 { //«Скорпіон»(23.10–22.11),
		fmt.Println("Скорпіон")
	}
	if m == 11 && d > 22 || m == 12 && d < 22 { //«Стрілець» (23.11–21.12),
		fmt.Println("Стрілець")
	}
	if m == 12 && d > 21 || m == 1 && d < 20 { //«Козеріг» (22.12–19.1).
		fmt.Println("Козеріг")
	}
}
