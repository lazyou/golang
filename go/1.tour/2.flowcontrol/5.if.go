package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	// 就像 `for` 循环一样，Go 的 `if` 语句也不要求用 `( )`将条件括起来，同时， `{ }` 还是必须有的。
	if x < 0 {
		return sqrt(-x) + "i"
	}

	// Sprint 使用其操作数的默认格式进行格式化并返回其结果字符串
	// 当两个连续的操作数均不为字符串时，它们之间就会添加空格
	return fmt.Sprint(math.Sqrt(x))
}

func main()  {
	fmt.Println(sqrt(2), sqrt(-4))
}