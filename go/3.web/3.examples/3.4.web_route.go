package main

import (
	"fmt"
	"net/http"
)

// 实现了一个简易的路由器
type MyMux struct {
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)

	if r.URL.Path == "/" {
		sayhelloName(w, r)
		return
	}

	http.NotFound(w, r)
	return
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello myroute!")
}

func main() {
	mux := &MyMux{}

	fmt.Println("localhost:9090 web start!")

	http.ListenAndServe(":9090", mux)
}
