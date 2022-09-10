package main

import (
	"fmt"
	"runtime"
)

// goroutine是通过Go的runtime管理的一个线程管理器。goroutine通过go关键字实现了，其实就是一个普通的函数。

// 下面的多个goroutine运行在同一个进程里面，共享内存数据，不过设计上我们要遵循：**不要通过共享来通信，而要通过通信来共享**
func say(s string) {
	for i := 0; i < 5; i++ {
		//开一个新的Goroutines执行
		// runtime.Gosched()表示让CPU把时间片让给别人,下次某个时候继续恢复执行该goroutine
		runtime.Gosched(); // sched 单词: 固定播送时间

		fmt.Println(s)
	}
}

func main() {
	// 通过关键字go就启动了一个goroutine
	go say("world")

	//当前Goroutines执行
	// TODO: 然而我还是看不懂为什么 先调用 say() 才能触发执行 go say(), 到底要如何理解 goroutine
	say("hello")
}