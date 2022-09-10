package main

import "fmt"

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

// 数值常量是高精度的 值
// 一个未指定类型的常量由上下文来决定其类型
func main() {
	// 21
	fmt.Println(needInt(Small))

	// 0.2
	fmt.Println(needFloat(Small))

	// 1.2676506002282295e+29
	fmt.Println(needFloat(Big))
}
