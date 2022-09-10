package main

import "fmt"

func main() {
	// channel 第二个参数指定了缓冲大小(非阻塞)
	c := make(chan int, 2) //修改2为1就报错，修改2为3可以正常运行

	c <- 1
	c <- 2

	fmt.Println(<-c)
	fmt.Println(<-c)
}