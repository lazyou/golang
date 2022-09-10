package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1

	for i := 0; i < n; i++ {
		c <- x

		x, y = y, x+y
	}

	// TODO: 如果不关闭 会产生一个 fatal error: all goroutines are asleep - deadlock!
	// 但是程序是如何知道需要 close 呢???
	close(c)
}

func main() {
	c := make(chan int, 10)

	go fibonacci(cap(c), c)

	for i := range c {
		fmt.Println(i)
	}
}
