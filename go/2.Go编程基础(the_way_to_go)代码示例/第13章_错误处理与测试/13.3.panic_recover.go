package main

import (
	"fmt"
)

func badCall() {
	panic("bad end")
}

func test() {
	// defer recover 一定要在这个位置
	// recover 只能在 defer 修饰的函数（参见 6.4 节）中使用：用于取得 panic 调用中传递过来的错误值，如果是正常执行，调用 recover 会返回 nil，且没有其它效果。
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("【2】Panicing %s\r\n", e)
		}
	}()

	badCall()
	fmt.Printf("【没到这里，进入了 recover()】After bad call\r\n")
}

func main() {
	fmt.Printf("【1】Calling test\r\n")
	test()
	fmt.Printf("【3】Test completed\r\n")
}
