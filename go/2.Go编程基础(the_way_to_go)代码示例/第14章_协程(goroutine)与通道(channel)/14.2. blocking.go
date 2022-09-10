package main

import (
	"fmt"
)

// 解释为什么下边这个程序会导致 panic：所有的协程都休眠了 - 死锁！
func main() {
	out := make(chan int)
	out <- 2
	go f1(out)
}

func f1(in chan int) {
	fmt.Println(<-in)
}
