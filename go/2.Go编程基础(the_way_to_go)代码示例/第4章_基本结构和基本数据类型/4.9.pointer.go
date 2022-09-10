package main

import (
	"fmt"
)

func main() {
	var i1 = 5
	fmt.Printf("value: %d; address: %p \n", i1, &i1)

	// 定义一个指针类型
	var intP *int
	intP = &i1
	fmt.Println(intP)
}
