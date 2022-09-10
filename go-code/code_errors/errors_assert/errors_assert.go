package main

import (
	"errors"
	"fmt"
)

type ApiError struct {
	Message string
}

func (a ApiError) Error() string {
	return a.Message
}

func getApiError() error {
	return ApiError{Message:"用户不存在"}
}

// 错误断言
func main() {
	err := getApiError()

	// 断言, 判断到类型
	if _, ok := err.(ApiError); ok {
		fmt.Printf("业务逻辑错误1: %s\n", err.Error())
	}

	// 不太好用, 错误信息必须一致, 使用 switch 可判断到类型
	if errors.Is(err, ApiError{Message:"用户不存在"}) {
		fmt.Printf("业务逻辑错误2: %s\n", err.Error())
	}
}
