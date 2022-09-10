package main

import "fmt"

// Go 的返回值可以被命名，并且就像在函数体开头声明的变量那样使用
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	// 返回值的命名在函数名上面声明了
	// 没有参数的 return 语句返回各个返回变量的当前值
	return
}

func main() {
	fmt.Println(split(17)) // 得到 "7 10"
}
