package main

import (
	"fmt"
	"math"
)

// Vertex 坐标
type Vertex struct {
	X, Y float64
}

// Go 没有类。不过你可以为结构体类型定义方法

// Abs TODO: 写法看起来更奇怪了. 函数名前面的括号放的是 结构体 (作为 接收者)
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
