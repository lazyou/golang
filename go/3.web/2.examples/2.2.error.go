package main

import (
	"errors"
	"fmt"
)

// 2.2 错误类型: TODO: 似乎没有介绍类似如何捕捉 错误 | 异常 的操作!
func main()  {
	err := errors.New("error 错误类型: 专门用来处理错误信息")

	if err !=nil {
		fmt.Println(err)
	}
}
