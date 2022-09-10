package main

import (
	"fmt"
)

func main() {
	// 空interface(interface{})不包含任何的method，正因为如此，所有的类型都实现了空interface
	var a interface{}

	var i int = 5

	s := "Hello world"

	// a可以存储任意类型的数值
	a = i
	a = s

	fmt.Println(a)
}
