package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// 注: 如果 func (v Vertex) Scale(f float64)  就不一样了 (去掉 * 号)
func (v *Vertex) Scale(f float64) {
	v.X = v.X *f
	v.Y = v.Y *f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

/*
接收者为指针的方法
	方法可以与命名类型或命名类型的指针关联。

	注意 (v *Vertex) 和 (v Vertex) 是有本质区别的
*/
func main()  {
	v := &Vertex{3, 4}

	// Before scaling: &{X:3 Y:4}, Abs: 5
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())

	// 结构体的值发生变化, 因为是指针
	v.Scale(5)

	// After scaling: &{X:15 Y:20}, Abs: 25
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())	
}