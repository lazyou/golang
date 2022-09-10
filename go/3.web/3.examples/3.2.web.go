package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// 全局变量, 不同浏览器打开也会不断累加 -- 生命周期与 脚本语言的不同之处
var globalCount int = 0

// http包建立Web服务器
func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fmt.Println(r.Form)
	// TODO: 为什么默认会有 /favicon.ico 这个请求?
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	// 全局
	globalCount++
	fmt.Println(globalCount)

	localCount := 1
	localCount++
	fmt.Println(localCount)

	//这个写入到 w 的是输出到客户端的
	fmt.Fprintf(w, "Hello astaxie!")
	fmt.Fprintf(w, "Hello astaxie!")
	fmt.Println()
}

func main() {
	http.HandleFunc("/", sayhelloName)

	fmt.Println("localhost:9090 web start!")

	err := http.ListenAndServe(":9090", nil)

	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
