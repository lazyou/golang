package main

import "fmt"

func test(c chan int) {
	for i := 0; i < 10; i++ {
		fmt.Println("send ", i)
		c <- i
	}
}

func main() {
	ch := make(chan int)

	go test(ch)

	for j := 0; j < 10; j++ {
		fmt.Println("get ", <-ch)
	}

	/**
	TODO: 为什么不是 一个 send 一个 get???
	send  0
	send  1
	get  0
	get  1
	send  2
	send  3
	get  2
	get  3
	send  4
	send  5
	get  4
	get  5
	send  6
	send  7
	get  6
	get  7
	send  8
	send  9
	get  8
	get  9
	*/
}
