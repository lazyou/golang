package main

import (
	"fmt"
)

// 函数 adder 返回一个闭包。每个闭包都被绑定在其各自的 sum 变量上
func adder() func(int) int {
	sum := 0

	// 匿名函数里使用了函数体外的变量, 使得 函数 被绑定在这个变量上, 从而形成闭包?
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()

	for i := 0; i < 3; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
