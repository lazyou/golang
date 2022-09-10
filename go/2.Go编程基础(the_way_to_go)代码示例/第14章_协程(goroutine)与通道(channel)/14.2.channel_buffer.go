package main

import "fmt"
import "time"

// 带缓冲的通道
func main() {
	// c := make(chan int) // TODO: 去掉注释对比
	c := make(chan int, 10)

	go func() {
		time.Sleep(15 * 1e9)
		x := <-c
		fmt.Println("received", x)
	}()

	fmt.Println("sending", 10)
	c <- 10
	fmt.Println("sent", 10)
}
