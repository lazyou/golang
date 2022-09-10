package main

import (
	"fmt"
	"time"
)

// 通道、超时和计时器（Ticker）
func main() {
	tick := time.Tick(1e9)
	boom := time.After(3e9)

	for {
		select {
		case <-tick:
			fmt.Println("tick. __每__ 1秒后触发!")
		case <-boom:
			fmt.Println("BOOM! 3秒后触发!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(1e9)
		}
	}
}