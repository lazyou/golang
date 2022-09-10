package main

import (
	"os"
	"net/http"
	"fmt"
	"io/ioutil"
)

// 发起一个 http 请求
// `go run 1.5.1.fetch.go http://www.php.test/`
func main() {
	for _, url := range os.Args[1:]{
		resp, err := http.Get(url)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		// resp 的 Body 字段包括一个可读的服务器响应流
		// ioutil.ReadAll 函数从 response 中读取到全部内容；将其结果保存在变量b中。
		body, err := ioutil.ReadAll(resp.Body)
		// resp.Body.Close 关闭 resp 的 Body 流，防止资源泄露
		resp.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

		fmt.Printf("%s", body)
	}
}