package main

import (
	"errors"
	"fmt"
)

var errWarn = errors.New("警告错误")
var errApi = errors.New("业务错误")

func getErrWarn() error {
	return errWarn
}

func getErrApi() error {
	return errApi
}

func main() {
	err := getErrApi()
	if errors.Is(err, errApi)  {
		// TODO: 但是错误内容却是固定的, 可以优化?
		fmt.Println("业务错误相关返回, 但是错误内容却是固定的, 并不实用, 比如: 要实现 数据不存在, 用户不存在 等错误信息提醒")
	}
}
