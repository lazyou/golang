package main

import (
	"fmt"
	"sync"
)

func cal(a int, b int, n *sync.WaitGroup) {
	c := a + b

	fmt.Printf("%d + %d = %d\n", a, b, c)

	defer n.Done() // goroutinue 完成后, WaitGroup 的计数 -1
}

func main() {
	var go_sync sync.WaitGroup // 声明一个 WaitGroup 变量

	for i := 0; i < 10; i++ {
		go_sync.Add(1) // WaitGroup 的计数加 1

		go cal(i, i+1, &go_sync)
	}

	go_sync.Wait() // 等待所有 goroutine 执行完毕 (TODO: 这里就不用再用什么 sleep 延长主线程的执行时间了)

	/**
	9 + 10 = 19
	0 + 1 = 1
	7 + 8 = 15
	8 + 9 = 17
	4 + 5 = 9
	6 + 7 = 13
	1 + 2 = 3
	2 + 3 = 5
	3 + 4 = 7
	5 + 6 = 11
	*/
}
