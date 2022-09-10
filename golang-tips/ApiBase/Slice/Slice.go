package main

import "fmt"

func main() {
	// 不等式永远成立：0 <= len(s) <= cap(s)
	// 使用内置函数make()可以给切片初始化, 指定切片类型和指定长度和可选容量的参数
	var slice = make([]int, 10)
	slice = []int{2, 3, 5, 7, 11}

	fmt.Printf("%v \n", slice)
	fmt.Printf("%#v \n", slice)
}
