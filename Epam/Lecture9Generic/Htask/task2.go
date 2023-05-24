package main

import "fmt"

type Map[K comparable, V any] struct {
	data map[K]V
}

func NewMap[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{
		data: make(map[K]V),
	}
}

func (m *Map[K, V]) Add(key K, value V) *Map[K, V] {
	m.data[key] = value
	return m
}

func (m *Map[K, V]) Delete(key K) {
	delete(m.data, key)
}

func (m *Map[K, V]) Update(key K, value V) {
	m.data[key] = value
}

func (m *Map[K, V]) Get(key K) (V, bool) {
	value, ok := m.data[key]
	return value, ok
}

func main() {
	myMap := NewMap[int, string]()                      // Создание экземпляра Map с ключами типа int и значениями типа string
	myMap.Add(1, "first").Add(3, "three").Add(2, "Two") // Добавление элемента в карту
	fmt.Println(myMap.data)                             // Вывод карты

	value, ok := myMap.Get(1) // Получение значения элемента по ключу
	if ok {
		fmt.Println("Value:", value)
	} else {
		fmt.Println("Key not found")
	}

	myMap.Update(1, "updated") // Обновление значения элемента
	fmt.Println(myMap.data)

	myMap.Delete(1) // Удаление элемента по ключу
	fmt.Println(myMap.data)
}
