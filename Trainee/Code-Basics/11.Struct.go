/*
Задание
На сервер приходит HTTP-запрос. Тело запроса парсится
и мапится в модель. Сразу работать с моделью нельзя,
потому что данные могут быть неверными.
Реализуйте функцию Validate(req UserCreateRequest) string,
которая валидирует запрос и возвращает строку
с ошибкой "invalid request", если модель невалидна.
Если запрос корректный, то функция возвращает пустую строку.
Правила валидации и структура модели описаны ниже.
Не используйте готовые библиотеки и опишите
правила самостоятельно.
*/
package main

import (
	"fmt"
	"strings"
)

// UserCreateRequest is a request to create a new user.
type UserCreateRequest struct {
	FirstName string
	Age       int
}

// BEGIN (write your solution here)
func Validate(req UserCreateRequest) string {
	if req.Age < 1 || req.Age > 150 {
		return "invalid request"
	}
	if strings.Contains(req.FirstName, " ") || req.FirstName == "" {
		return "invalid request"
	}
	return "ok"
}

func main() {
	user := UserCreateRequest{"proverka", 35}
	user1 := UserCreateRequest{"prov erka", 35}
	user2 := UserCreateRequest{"proverka", 235}
	user3 := UserCreateRequest{"", 35}
	user4 := UserCreateRequest{"proverka", 0}
	fmt.Println(Validate(user))
	fmt.Println(Validate(user1))
	fmt.Println(Validate(user2))
	fmt.Println(Validate(user3))
	fmt.Println(Validate(user4))
}

// END
