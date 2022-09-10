package main

import "fmt"

// new() 函数: 返回的值是指向新值的指针, 分配了该类型的零值
func main() {
	// new 函数返回指针
	p := new(int)

	// 打印指针(内存地址) 0xc0420080c0
	fmt.Printf("%v \n", p)

	// 打印指针指向的值
	fmt.Printf("%v \n", *p)

	*p = 2          // 设置 int 匿名变量的值为 2
	fmt.Println(*p) // "2"
}
