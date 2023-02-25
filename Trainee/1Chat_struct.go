/*
Напишите программу для учета продуктов в магазине.

	Создайте структуру "Продукт" (Product) с полями
	 "название" (Name), "цена" (Price) и
	  "количество" (Quantity).

	 Затем создайте несколько экземпляров
	 структуры "Продукт" с различными названиями,
	 ценами и количествами. Создайте функцию,
	 которая принимает на вход массив продуктов
	 и выводит на экран информацию о каждом продукте,
	а также общую стоимость всех продуктов в магазине.
*/
package main

import "fmt"

type Produkt struct {
	Name     string
	Price    float64
	Quantity float64
}

func kolTovar(poziciya []Produkt) {
	obsh_Kol := 0.0
	for _, v := range poziciya {

		fmt.Printf("Наименование: %s Кол-во: %.2f, Цена за 1eд: $(%.2f)\n",
			v.Name, v.Quantity, v.Price)

		obsh_Kol = v.Quantity * v.Price
		fmt.Printf("%s есть на сумму $%.2f\n", v.Name, obsh_Kol)
	}
}

func main() {
	tovar := []Produkt{
		{"orange", 56, 7.5},
		{"beer", 20, 56},
		{"milk", 33.5, 70},
	}
	kolTovar(tovar)
}
