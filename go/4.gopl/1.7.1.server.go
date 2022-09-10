package main

import (
	"net/http"
	"log"
	"fmt"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// 把结果发送到响应中，这里我们用的是标准输出流的fmt.Fprintf
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
