package main

import (
	"fmt"
)

type ApiError struct {
	Message string
}

func (a *ApiError) Error() string {
	return a.Message
}

func getApiError() error {
	return &ApiError{Message:"用户不存在"}
}

// 错误断言: 可判断到类型, 不必错误信息都一致
func main() {
	err := getApiError()

	switch e := err.(type) {
	case *ApiError:
		fmt.Println("ApiErr:")
		fmt.Println(e)
	default:
		fmt.Printf("Not a special error, just %s\n", err)
	}
}
