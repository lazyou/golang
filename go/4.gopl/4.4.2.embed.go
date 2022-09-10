package main

import "fmt"

type Point struct {
	X, Y int
}

// 结构体嵌套
type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	w1 := Wheel{Circle{Point{8, 8}, 5}, 20}

	// 推荐声明方式
	w2 := Wheel{
		Circle: Circle{
			Point: Point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20,
	}

	fmt.Printf("%#v\n", w1)
	fmt.Printf("%#v\n", w2)
}