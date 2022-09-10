package main

import (
	"fmt"
)

func main() {
    c := make(chan int)

	go func() {
		fmt.Println("goroutine message")
		c <- 1 // 步骤1
    }()
    
    i := <-c // 步骤2
	fmt.Println(i) // 输出 1
    
	fmt.Println("main function message")
}
