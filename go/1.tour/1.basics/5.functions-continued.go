package main

import "fmt"

func main() {
	fmt.Println(add(42, 13))
}

// 当两个或多个连续的函数命名参数是同一类型，则除了最后一个类型之外，其他都可以省略
func add(x, y int) int {
	return x + y
}
