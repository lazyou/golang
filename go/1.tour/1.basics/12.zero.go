package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string

	// 变量在定义时 **没有明确的初始化** 时会赋值为 **零值**
	// 0 0 false ""
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
