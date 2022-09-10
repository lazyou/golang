package main

import (
	"fmt"
	"time"
)

// 从不同的并发执行的协程中获取值可以通过关键字select来完成，它和switch控制语句非常相似（章节5.3）也被称作通信开关；
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go pump1(ch1)
	go pump2(ch2)
	go suck(ch1, ch2)

	time.Sleep(1e9)
}

func pump1(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i * 2
	}
}

func pump2(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i + 5
	}
}

func suck(ch1, ch2 chan int) {
	for {
		select {
		case v := <-ch1:
			fmt.Printf("Received on channel 1: %d\n", v)
		case v := <-ch2:
			fmt.Printf("Received on channel 2: %d\n", v)
		}
	}
}
