package main

import "fmt"

func main() {
	fmt.Println(add(42, 13))
}

// 函数可以没有参数或接受多个参数
// 这里, add 接受两个 int 类型的参数
func add(x int, y int) int {
	return x + y
}
