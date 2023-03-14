// Рассмотрим массив/список овец, где некоторые овцы
// могут отсутствовать на своем месте. Нам нужна функция,
// которая подсчитывает количество овец, присутствующих
// в массиве (true означает наличие).
package main

import "fmt"

func main() {
	var arr = [...]bool{true, true, true, false,
		true, true, true, true,
		true, false, true, false,
		true, false, false, true,
		true, true, true, true,
		false, false, true, true}
	var count int
	for _, v := range arr {
		if v == true {
			count += 1
		}
	}
	fmt.Println(count)
}

// package kata

// func CountSheeps(numbers []bool) int {
//   count:=0
//   for _,v:=range numbers{
//     if v== true{
//       count+=1
//     }
//   }
//   return count // your code here
// }
