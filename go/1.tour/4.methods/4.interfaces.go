package main

import (
	"fmt"
	"math"
)

// 接口: 接口类型是由一组方法定义的集合
// 接口类型的值可以存放实现这些方法的任何值
type Abser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}

	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func main()  {
	// 接口
	var a Abser

	// Abs() 方法
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	// 把 Abs() 方法的实现赋予 接口 a
	a = f
	a = &v // 指针类型的赋值方式要加 & 符号?

	// a = v
	fmt.Println(a.Abs())
}
