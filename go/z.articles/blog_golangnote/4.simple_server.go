package main

import (
	"log"
	"html"
	"fmt"
	"net/http"
)

func main()  {
    // 1. TODO: 这个 `http.Handle` 要怎么使用
    // http.Handle("/foo", fooHandler)

    // 2.
    http.HandleFunc("/bar", func (w http.ResponseWriter, r *http.Request)  {
        fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
    })

    log.Fatal(http.ListenAndServe(":8000", nil))
}
