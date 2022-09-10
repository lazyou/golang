package main

import (
	"fmt"
)

type Rectangle struct {
	width, height float64
}

// (r Rectangle) 这里不要理解成接收的参数, 更像函数的类型
func (r Rectangle) area() float64 {
	return r.width * r.height
}

func main() {
	r := Rectangle{12, 2}

	fmt.Println(r.area())
}
