package main

import (
	"fmt"
	"time"
)

func main() {
	// time.Tick() 函数声明为 Tick(d Duration) <-chan Time
	// 当你想返回一个通道而不必关闭它的时候这个函数非常有用：它以 d 为周期给返回的通道发送时间，d是纳秒数。
	tick := time.Tick(1e8)

	// 可以在 select 中通过 time.After() 发送的超时信号来停止协程的执行
	boom := time.After(3e8)

	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(1e8)
		}
	}
}
