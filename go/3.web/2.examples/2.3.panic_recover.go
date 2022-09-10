package main

import "fmt"

// panic Go没有像Java那样的异常机制，它不能抛出异常，而是使用了panic和recover机制。
// __一定要记住，你应当把它作为最后的手段来使用__
func main() {
	result := throwsPanic(productPanic)

	fmt.Println(result)
}

func throwsPanic(f func()) (b bool) {
	defer func() {
		if x := recover(); x != nil {
			b = true
		}
	}()

	f() // 执行函数 f，如果f中出现了 panic，那么就可以恢复回来. TODO: 没理解所谓的恢复回来是为了做什么?

	return
}

func productPanic() {
	panic("还是不懂 go 的 panic 机制!")

	fmt.Println("panic 不中断")
}
