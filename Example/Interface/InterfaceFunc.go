package main

import "fmt"

//Создаем интерфейс
type greeter interface {
	greet(string) string
}

//Создаем структуры
type ukrainian struct{}
type english struct{}

//Создаем методы для структур
func (u *ukrainian) greet(name string) string {
	return fmt.Sprintf("Привіт, %s", name)
}
func (e *english) greet(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}

//передаем в вункцию интерф  и переменную
func sayHallo(g greeter, name string) {
	fmt.Println(g.greet(name)) // вызываем метод для name
}

func main() {
	// Вызов функий(адрес струткуры)
	sayHallo(&ukrainian{}, "Микола")
	sayHallo(&english{}, "John")
}
