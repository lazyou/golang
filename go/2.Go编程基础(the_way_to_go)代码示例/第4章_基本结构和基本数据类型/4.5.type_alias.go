package main

import "fmt"

type INT_ALIAS int

// 类型别名
func main() {
	var a, b INT_ALIAS = 3, 4
	c := a + b
	fmt.Printf("c has the value: %d", c)
}
