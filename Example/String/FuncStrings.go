package main

import (
	"fmt"
	"strings"
)

func main() {
	//логическая проверка есть ли в тексте "abs"
	fmt.Println(strings.Contains("Test Text", "abs"))
	//любой символ из "abs"
	fmt.Println(strings.ContainsAny("Test Text", "abs"))
	//солько раз в наш текст входит "ana" без нахлестов букв
	fmt.Println(strings.Count("Banana", "ana"))
	//Есть ли "Tes" префиксом рядка
	fmt.Println(strings.HasPrefix("Test Text", "Tes"))
	//Есть ли "es" суфиксом рядка
	fmt.Println(strings.HasSuffix("Test Text", "es"))
	//выводит индекс с какой позиции начинается "ext"
	fmt.Println(strings.Index("Test Text", "ext"))
	//выводит индекс любого из подходящих символов списка "abt"
	fmt.Println(strings.IndexAny("Test Text", "abt"))
	//выводит индекс из последних символов списка "abt"
	fmt.Println(strings.LastIndex("Test Text", "abc"))
	//выводит индекс из последнего набора символов. списка "abt"
	fmt.Println(strings.LastIndexAny("Test Text", "abt"))
	// буквы "о" заменили на "." дважды
	fmt.Println(strings.Replace("foo", "o", ".", 2))

	//сдвиг символов на одну позицию в право по коду Аски.
	// Было АБ стало БВ
	f := func(r rune) rune {
		return r + 1
	}
	fmt.Println(strings.Map(f, "ab"))

	//все буквы верхний регистр
	fmt.Println(strings.ToUpper("Test Text"))
	//все буквы нижний регистр
	fmt.Println(strings.ToLower("Test Text"))
	//первые буквы слов всего текста вверх
	fmt.Println(strings.Title("test text"))
	//удаление пробелов и экскейп последовательностей
	fmt.Println(strings.TrimSpace(" foo\n"))
	//вырезает указанные символы "fo"
	fmt.Println(strings.Trim("fooe", "fo"))
	//вырезает указанные символы слева
	fmt.Println(strings.TrimLeft("foo", "f"))
	//вырезает указанные символы правые,все
	fmt.Println(strings.TrimRight("foo", "o"))
	//вырезает префикс "fo"
	fmt.Println(strings.TrimPrefix("foo", "fo"))
	//вырезает суфикс "o"
	fmt.Println(strings.TrimSuffix("foo", "o"))
	//убирает все символы пробела и табуляции
	//результатом функции есть массив, не рядок!
	fmt.Println(strings.Fields(" a\t b\n"))
	//разделить рядок с помощью "," на элементы массива
	// "," НЕ СОХРАНЯЕТСЯ
	fmt.Println(strings.Split("a,b", ","))
	//разделить рядок с помощью "," на элементы массива
	// "," СОХРАНЯЕТСЯ
	fmt.Println(strings.SplitAfter("a,b", ","))
	//функция обьединяет массив в рядок разделяя символы ":"
	fmt.Println(strings.Join([]string{"a", "b"}, ":"))
	//Репит двойка это клоичество повторений рядка
	fmt.Println(strings.Repeat("Epam", 2))

}
