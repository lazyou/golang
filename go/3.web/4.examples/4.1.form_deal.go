package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
)

func main() {
	http.HandleFunc("/login", login) // 路由

	fmt.Println("localhost:9090 start!")

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)

	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		// 请求的数据
		fmt.Println("username:", r.FormValue("username"))
		fmt.Println("password:", r.FormValue("password"))

		// request.Form 是一个 url.Values 类型，里面存储的是对应的类似key=value的信息
		v := url.Values{}
		v.Set("name", "Ava")
		v.Add("friend", "Jess")
		v.Add("friend", "Sarah")

		// v.Encode() == "name=Ava&friend=Jess&friend=Sarah&friend=Zoe"
		fmt.Println(v.Get("name"))
		fmt.Println(v.Get("friend"))
		fmt.Println(v["friend"])
		fmt.Println(v)
		fmt.Println() // map[friend:[Jess Sarah] name:[Ava]]

		// 转义
		fmt.Println("username", template.HTMLEscapeString(r.FormValue("username")))
		template.HTMLEscape(w, []byte(r.Form.Get("username"))) //输出到客户端
	}
}
