package main

import (
	"fmt"
)

// 斐波纳契（一种整数数列）
func fibonacci(n int, c chan int) {
	x, y := 1, 1

	for i := 0; i < n; i++ {
		c <- x

		x, y = y, x + y
	}

	// 生产者通过内置函数close关闭channel. (PS: 在消费方可以通过语法v, ok := <-ch测试channel是否被关闭)
	// 另外记住一点的就是channel不像文件之类的，不需要经常去关闭，只有当你确实没有任何发送数据了，或者你想显式的结束range循环之类的
	close(c)
}

func main() {
	c := make(chan int, 10)

	fmt.Println(c)	// 0xc820062000 内存地址
	fmt.Println(cap(c)) // 10

	go fibonacci(cap(c), c)

	// for i := range c能够不断的读取channel里面的数据，直到该channel被显式的关闭
	for i := range c {
		fmt.Println(i)
	}
}