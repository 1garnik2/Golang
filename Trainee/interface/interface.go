package main

import "fmt"

type greeter interface {
	greet(string) string
}

type uk struct{}
type en struct{}

func (u *uk) greet(name string) string {
	return fmt.Sprintf("Privit %s!", name)
}
func (e *en) greet(name string) string {
	return fmt.Sprintf("Hello %s!", name)
}

func sayHello(g greeter, name string) {
	fmt.Println(g.greet(name))
}

func main() {
	sayHello(&uk{}, "Petr")
	sayHello(&en{}, "Djoe")
}
