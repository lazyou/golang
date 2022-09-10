package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// // 指针接收者 -- 使得内部的值被修改了
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Vertex) ScaleWithoutPointer(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	// 案例: 1
	vPointer := Vertex{3, 4}
	// 5
	fmt.Println(vPointer.Abs())

	// 指针接收者 -- 影响结构体内部字段
	vPointer.Scale(10)
	// 50
	fmt.Println(vPointer.Abs())

	// 案例: 2
	vNoPointer := Vertex{3, 4}
	// 5
	fmt.Println(vNoPointer.Abs())

	// 非指针接收者 -- 不影响结构体内部字段
	vNoPointer.ScaleWithoutPointer(10)
	// 5
	fmt.Println(vNoPointer.Abs())
}
