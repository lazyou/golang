package main

import "fmt"

func main() {
	fmt.Println("counting")

	// defer 栈: **后进先处** 的顺序调用被延迟的函数调用
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
	/*
		counting
		done
		9
		8
		7
		6
		5
		4
		3
		2
		1
		0
	*/
}
