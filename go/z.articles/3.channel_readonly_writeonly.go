package main

import (
	"fmt"
	"time"
)

// 只能向chan里写数据
func send(c chan<- int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
}

// 只能取channel中的数据
func get(c <-chan int) {
	for i := range c {
		fmt.Println(i)
	}
}

func main() {
	c := make(chan int)

	go send(c)
	go get(c)

	time.Sleep(time.Second * 1)

	/**
	0
	1
	2
	3
	4
	5
	6
	7
	8
	9
	*/
}
