package main

import (
	"net/http"
	"log"
	"sync"
	"fmt"
)

var mu sync.Mutex
var count int

func main()  {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	printMore(w, r)

	// 并发情况下，假如真的有两个请求同一时刻去更新count，那么这个值可能并不会被正确地增加；这个程序可能会引发一个严重的bug：竞态条件

	// 为了避免这个问题，我们必须保证每次修改变量的最多只能有一个 goroutine，这也就是代码里的 mu.Lock() 和 mu.Unlock() 调用将修改 count 的所有行为包在中间的目的
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.PATH = %q\n", r.URL.Path)
}

// 响应次数统计
func counter(w http.ResponseWriter, r *http.Request) {
	printMore(w, r)

	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	fmt.Println(count)
	mu.Unlock()
}

// 把请求的http头和请求的form数据都打印出来
func printMore(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)

	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}

	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}