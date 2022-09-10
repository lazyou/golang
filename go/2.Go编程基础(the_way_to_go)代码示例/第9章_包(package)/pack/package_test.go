package main

import (
	"fmt"

	"./pack1"
)

// 然而只有本目录在 gopath 下才能引用相对路径的包
// 否则是找不到的
func main() {
	var test1 string
	test1 = pack1.ReturnStr()
	fmt.Printf("ReturnStr from package1: %s\n", test1)
	fmt.Printf("Integer from package1: %d\n", pack1.Pack1Int)
	// fmt.Printf("Float from package1: %f\n", pack1.pack1Float)
}
