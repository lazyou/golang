package main

import (
	"fmt"
)

func main() {
	fmt.Println("counting")

	// defer 栈: 先进后出 (也就是 后进先出, 所以从 9 开始打印到 1)
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
