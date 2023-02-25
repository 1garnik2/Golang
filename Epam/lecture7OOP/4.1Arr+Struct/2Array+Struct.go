package main

import "fmt"

type book struct {
	title        string
	Author       string
	YearReleased int
	Publisher    string
	NumPages     int
}

func aboutBook(s book) {
	fmt.Printf("Название : %s\n", s.title)
	fmt.Printf("Автор : %s\n", s.Author)
	fmt.Printf("Год релиза : %d \n", s.YearReleased)
	fmt.Printf("Издание : %s\n", s.Publisher)
	fmt.Printf("Количество страниц : %d\n", s.NumPages)
}
func main() {
	book1 := [2]book{
		{"Book", "Bob", 1999, "typograf", 678},
		{"Book1", "Rob", 1969, "Typo", 68},
	}

	for i := 0; i < len(book1); i++ {
		if book1[i].title == "Book" {
			aboutBook(book1[i])
			fmt.Println()
		}
	}

}
