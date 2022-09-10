package main

import (
	"fmt"
	"time"
)

func doSomethingElseForAWhile() int {
	fmt.Println("doing.")
	time.Sleep(1e9)
	fmt.Println("done.")
	return 99
}

func compute(ch chan int){
	ch <- 1 // when it completes, signal on the channel.
}

// 信号量模式
// 协程通过在通道 ch 中放置一个值来处理结束的信号。main 协程等待 <-ch 直到从中获取到值。
func main(){
	ch := make(chan int) 	// allocate a channel.

	go compute(ch)		// start something in a goroutines
	doSomethingElseForAWhile()
	result := <- ch
	fmt.Println(result)
}
