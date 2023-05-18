package main

import "fmt"

type Point struct {
	X int
	Y int
}

func main() {
	pointsMap := map[string]Point{}
	otherPointMap := map[string]Point{}
	otherPointMap["a"] = Point{5, 7}
	fmt.Println(otherPointMap)
	pointsMap["a"] = Point{1, 2}
	pointsMap["b"] = Point{4, 6}
	fmt.Println(pointsMap)
}
