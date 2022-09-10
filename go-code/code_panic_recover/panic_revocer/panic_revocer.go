package main

import (
	"fmt"
)

// （recover）内建函数被用于从 panic 或 错误场景中恢复：让程序可以从 panicking 重新获得控制权，停止终止过程进而恢复正常执行
// 总结: panic 会导致栈被展开直到 defer 修饰的 recover() 被调用或者程序中止

func badCall() {
	panic("bad end")
}

func test() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("Panicing %s\r\n", e)
		}
	}()

	badCall()

	fmt.Printf("After bad call\r\n") // 因为 panic 不会执行到这
}

func main() {
	fmt.Printf("Calling test\r\n")
	test()
	fmt.Printf("Test completed\r\n")
}
