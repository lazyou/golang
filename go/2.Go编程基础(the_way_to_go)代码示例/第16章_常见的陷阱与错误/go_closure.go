package main

import (
	"fmt"
	"time"
)

var values = [5]int{10, 11, 12, 13, 14}

func main() {
	// 版本A:
	for index := range values { // index是索引值
		func() {
			fmt.Print(index, " ")
		}() // 调用闭包打印每个索引值
	}
	fmt.Println("版本A end")

	// 版本B: 和A版本类似，但是通过调用闭包作为一个协程
	for index := range values {
		go func() {
			fmt.Print(index, " ")
		}()
	}
	fmt.Println("版本B end")
	time.Sleep(5e9)

	// 版本C: 正确的处理方式
	for index := range values {
		go func(index interface{}) {
			fmt.Print(index, " ")
		}(index)
	}
	fmt.Println("版本C end")
	time.Sleep(5e9)

	// 版本D: 输出值:
	for index := range values {
		val := values[index]
		go func() {
			fmt.Print(val, " ")
		}()
	}
	fmt.Println("版本D end")
	time.Sleep(1e9)
}
