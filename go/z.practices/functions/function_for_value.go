package main

import (
	"fmt"
	"math"
)

// TODO: 函数接收的参数是函数, 看起来真蛋疼
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	// 函数赋值给变量
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(5, 12))

	// 变量函数当做参数传递
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}
