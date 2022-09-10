package main

import "fmt"

func main() {
	// Go 语言中的数组是一种值类型（不像 C/C++ 中是指向首元素的指针）
	// 所以可以通过 new() 来创建 ( new 创建的是数组指针)
	var arr = new([5]int)
	arrClone := arr

	arr[0] = 1
	fmt.Printf("%v \n", arr)
	fmt.Printf("%v \n", arrClone)

	/**
	&[1 0 0 0 0]
	&[1 0 0 0 0]
	*/
}
