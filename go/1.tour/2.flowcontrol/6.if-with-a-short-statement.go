package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	// if 的便捷语句: 跟 for 一样， if 语句可以在条件之前执行一个简单语句。
	// 由这个语句定义的变量的作用域仅在 if 范围之内
	// 说实话, 这写法看着蛋疼
	if v := math.Pow(x, n); v < lim {
		return v
	}

	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}