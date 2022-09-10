package main

import "fmt"

// 使用 defer 语句实现代码追踪
// 然而实际上多个 defer 的各种嵌套看起来执行流程并不直观
func main() {
	b()
}

func b() {
	trace("b")
	defer untrace("b")
	fmt.Println("in b")
	a()
}

func a() {
	trace("a")
	defer untrace("a")
	fmt.Println("in a")
}

func trace(s string)   { fmt.Println("entering:", s) }
func untrace(s string) { fmt.Println("leaving:", s) }
