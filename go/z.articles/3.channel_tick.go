package main

import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan int, 5)

	for i := 1; i < 5; i++ {
		requests <- i
	}

	close(requests)

	limiter := time.Tick(time.Second * 1)

	for req := range requests {
		<-limiter

		fmt.Println("requests", req, time.Now()) // 执行到这里，需要隔1秒才继续往下执行，time.Tick(timer)上面已定义
	}

	/**
	requests 1 2018-07-16 00:16:50.113232294 +0800 CST m=+1.000652185
	requests 2 2018-07-16 00:16:51.113253198 +0800 CST m=+2.000672994
	requests 3 2018-07-16 00:16:52.113217549 +0800 CST m=+3.000637354
	requests 4 2018-07-16 00:16:53.11325736 +0800 CST m=+4.000677557
	*/
}
