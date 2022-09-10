package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// 定义一个类型为  (v *Vertex) 的 函数 Abs?
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func main()  {
	v := &Vertex{3, 4}

	// **Go 没有类**。然而，仍然可以在结构体类型上定义方法
	// 调用结构体上的方法
	fmt.Println(v.Abs())
}