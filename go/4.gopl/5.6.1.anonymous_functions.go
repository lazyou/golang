package main

import "fmt"

// 函数squares返回另一个类型为 func() int 的函数
// 该匿名函数每次被调用时都会返回下一个数的平方
func squares() func() int {
	var x int

	// 变量的生命周期不由它的作用域决定：squares返回后，变量x仍然隐式的存在于匿名函数中
	return func() int {
		x++
		return x * x
	}
}

func main() {
	f := squares()

	// squares的例子证明，函数值不仅仅是一串代码，还记录了状态。
	fmt.Println(f()) // "1"
	fmt.Println(f()) // "4"
	fmt.Println(f()) // "9"
	fmt.Println(f()) // "16"
}