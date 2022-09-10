// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 20.
//!+

// Server2 is a minimal "echo" and counter server.
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

// 服务器每一次接收请求处理时都会另起一个goroutine，这样服务器就可以同一时间处理多个请求。
// 然而在 __并发__ 情况下，假如真的有两个请求同一时刻去更新 count，那么这个值可能并不会被正确地增加；这个程序可能会引发一个严重的bug： __竞态条件__（参见9.1）

// 为了避免这个问题，我们必须保证每次修改变量的最多只能有一个 goroutine，这也就是代码里的 `mu.Lock()` 和 `mu.Unlock()` 调用将修改 `count` 的所有行为包在中间的目的。第九章中我们会进一步讲解 __共享变量__。

// 如何验证： mu.Lock() mu.Unlock() 的效果
// ab -n 1000 -c 100 http://localhost:8000/ 有用 Lock Unlock 时计数正常
// ab -n 1000 -c 100 http://localhost:8000/ 注释掉 Lock Unlock 时计数不正常
func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock() // 注释掉， 并用 ab 测试
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	fmt.Println(count)
}

// counter echoes the number of calls so far.
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	fmt.Println(count)
	mu.Unlock()
}

//!-
