package main

import (
	"fmt"
	"math"
)

// 数也是值。他们可以像其他值一样传递，比如，函数值可以作为函数的参数或者返回值

// compute函数 接收一个函数 (但是要定义函数的参数 返回值, 有点麻烦)
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)	
}

func main()  {
	hypot := func (x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	// 13
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	
	fmt.Println(compute(math.Pow))
}