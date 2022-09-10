package main

import (
	"fmt"
)

type Request struct {
	a, b      int
	replyChan chan int // reply channel inside the Request
}

type binOp func(a, b int) int

func run(op binOp, req *Request) {
	req.replyChan <- op(req.a, req.b)
}

func server(op binOp, service chan *Request) {
	for {
		req := <-service // requests arrive here
		// start goroutine for request:
		go run(op, req) // don't wait for op
	}
}

func startServer(op binOp) chan *Request {
	reqChan := make(chan *Request)
	go server(op, reqChan)
	return reqChan
}

// TODO: 没看懂
func main() {
	handle := func(a, b int) int { return a + b }
	adder := startServer(handle)

	// 创建 N 个 Request
	const N = 100
	var reqs [N]Request
	for i := 0; i < N; i++ {
		req := &reqs[i]
		req.a = i
		req.b = i + N
		req.replyChan = make(chan int)

		// TODO: 这部干啥?
		adder <- req
	}

	// checks:
	for i := N - 1; i >= 0; i-- { // doesn't matter what order
		replyChan := <-reqs[i].replyChan
		//fmt.Println("replyChan: ", replyChan)

		if replyChan != N+2*i {
			fmt.Println("fail at", i)
		} else {
			fmt.Println("Request ", i, " is ok!")
		}
	}

	fmt.Println("done")
}