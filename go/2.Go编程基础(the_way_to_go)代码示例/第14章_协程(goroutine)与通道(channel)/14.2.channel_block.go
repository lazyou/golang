package main

import "fmt"

// 14.2.3 通道阻塞
// 默认情况下，通信是同步且无缓冲的：在有接受者接收数据之前，发送不会结束
func main() {
	ch1 := make(chan int)
	go pump(ch1)       // pump hangs
	fmt.Println(<-ch1) // prints only 0
	fmt.Println(<-ch1) // prints only 1
	fmt.Println(<-ch1) // prints only 2
	// 说明 ch1 里面还有数据，只是因为进程结束了所以看不到
}

func pump(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
}
