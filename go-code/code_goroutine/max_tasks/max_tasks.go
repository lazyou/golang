package main

import (
	"fmt"
	"time"
)

const MAXREQS = 50

var sem = make(chan int, MAXREQS)
var count = 1

type Request struct {
	a, b   int
	replyc chan int
}

func process(r *Request) {
	count++
	fmt.Println(count)
}

func handle(r *Request) {
	sem <- 1 // doesn't matter what we put in it
	process(r)
	<-sem // one empty place in the buffer: the next request can start
}

func server(service chan *Request) {
	for {
		request := <-service
		go handle(request)
	}
}

// TODO: 又看不懂了...
// 利用 channels 限制同时处理的请求数
func main() {
	service := make(chan *Request)
	go server(service)
	time.Sleep(1e8)
}