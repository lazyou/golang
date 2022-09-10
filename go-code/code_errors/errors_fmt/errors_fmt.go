package main

import (
	"fmt"
)

// 用 fmt 创建错误对象
func fmtError() error {
	return fmt.Errorf("Err %s", "错误")
}

func main() {
	fmt.Println(fmtError())
}
