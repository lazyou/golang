package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	// if 的简短语句, if 后面跟随表达式, 表达式结果用来做 if 条件
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
