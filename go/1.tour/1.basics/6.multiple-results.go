package main

import "fmt"

// 多个返回值, 每个返回值的类型都要声明
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	// := 短变量 简洁赋值语句
	a, b := swap("hello", "world")

	fmt.Println(a, b)
}
