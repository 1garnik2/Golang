package main

import "fmt"

type Point struct {
	x int
	y int
}

func (p Point) method() {
	fmt.Println(p.y)
}

func main() {
	pointsMap := map[string]Point{}      // инициализация  map
	otherPointMap := make(map[int]Point) // инициализация  map без {}

	pointsMap["a"] = Point{
		x: 3,
		y: 6,
	}
	fmt.Println(pointsMap)
	fmt.Println(pointsMap["a"])
	otherPointMap[4] = Point{6, 8}
	otherPointMap[1] = Point{7, 9}
	otherPointMap[2] = Point{4, 6}
	otherPointMap[3] = Point{3, 56}
	fmt.Println(otherPointMap)
	fmt.Println(otherPointMap[1].x)
	otherPointMap[4].method()
	for k, v := range otherPointMap {
		fmt.Printf("Key Map = %d Value Map = %d\n", k, v)
	}

}
