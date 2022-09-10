package main

import "fmt"

/*
函数的闭包
	Go 函数可以是一个 **闭包**。
	
	**闭包是一个函数值，它引用了函数体之外的变量。**
	
	这个函数可以对这个引用的变量进行访问和赋值；换句话说这个函数被“绑定”在这个变量上。

例如，函数 adder 返回一个闭包。每个返回的闭包都被绑定到其各自的 sum 变量上。
*/

// 返回一个函数, 返回函数的返回值是 int
func adder() func(int) int {
	sum := 0
	
	return func (x int) int {
		sum += x

		return sum
	}
}

func main()  {
	pos, neg := adder(), adder()

	/*
	0 0
	1 -2
	3 -6
	6 -12
	10 -20
	15 -30
	21 -42
	28 -56
	36 -72
	45 -90	
	*/
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2 * i),
		)
	}
}