package main

import (
	"fmt"
)

// 通过类型别名的方式 -- 为非结构体类型声明方法 -- 不能为内建类型声明方法
type MyFloat float64

// 设置 别名类型 为接收者
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}

	return float64(f)
}

func main() {
	f := MyFloat(-100)

	fmt.Println(f.Abs())
}
