package main

import (
	"fmt"
	"time"
)

// 通道阻塞: 通信是同步且无缓冲的：在有接受者接收数据之前，发送不会结束
func main() {
	ch1 := make(chan int)

	go pump(ch1)       // pump hangs
	fmt.Println(<-ch1) // prints only 0. 只能接收到第一个 来自channel 的数据

	time.Sleep(1e9)
}

func pump(ch chan int) {
	for i := 10; ; i++ {
		ch <- i
	}
}
