package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	// 如果我们移除一个或所有 go 关键字，程序无法运行，Go 运行时会抛出 `panic`
	go sendData(ch)
	go getData(ch)

	// main() 等待了 1 秒让两个协程完成，如果不这样，sendData() 就没有机会输出。
	time.Sleep(1e9)
}

func sendData(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokyo"
}

// getData() 使用了无限循环：它随着 sendData() 的发送完成和 `ch` 变空也结束了。
func getData(ch chan string) {
	var input string
	// time.Sleep(2e9)
	for {
		input = <-ch
		fmt.Printf("%s ", input)
	}
}
