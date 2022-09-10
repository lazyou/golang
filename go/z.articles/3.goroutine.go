package main

import (
	"fmt"
	"runtime"
	"time"
)

func cal(a int, b int) {
	c := a + b
	fmt.Printf("%d + %d = %d \n", a, b, c)
}

func main() {
	num := runtime.NumCPU() // 获取主机的逻辑CPU个数

	fmt.Printf("主机的逻辑CPU个数为 %d 个\n", num)

	for i := 0; i < 10; i++ {
		// 每个 goroutine 的执行顺序是随机的
		go cal(i, i+1)
	}

	time.Sleep(time.Second * 2) // sleep 作用是为了等待所有任务完成
}
