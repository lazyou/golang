package main

import (
	"fmt"
	"time"
)

// 从不同的并发执行的协程中获取值可以通过关键字select来完成
// select 语句实现了一种监听模式，通常用在（无限）循环中；在某种情况下，通过 break 语句使循环退出。
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
		// 如果没有 default，select 就会一直阻塞。
		// 如果没有通道操作可以处理并且写了 default 语句，它就会执行：default 永远是可运行的（这就是准备好了，可以执行）
		//default:
		//	fmt.Println("default.")
		}
	}
}