package main

import (
	"fmt"
)

func add1(a *int) int {
	// 基本类型的指针类型在使用过程中必须加上 * 号
	*a = *a + 1

	return *a
}

func main() {
	x := 3
	fmt.Println(x)

	x1 := add1(&x)
	fmt.Println(&x)
	fmt.Println(x1)
	fmt.Println(x)
}
