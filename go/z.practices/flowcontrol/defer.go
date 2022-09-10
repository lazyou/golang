package main

import (
	"fmt"
)

func main() {
	// defer 语句会将函数推迟到外层函数返回之后执行
	defer fmt.Println("world")

	fmt.Println("hello")
}
