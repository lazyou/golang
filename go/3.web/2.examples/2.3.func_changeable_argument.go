package main

import "fmt"

// 变量 arg 是一个 int 的 slice
func myFunc(arg ...int) {
	for _, n := range arg {
		fmt.Println(n)
	}
}

func main() {
	myFunc(1, 2, 3)
}
