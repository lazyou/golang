package main

import (
	"fmt"
	"math"
)

// 定义了 I 接口
type I interface {
	M()
}

type T struct {
	S string
}

// 这里是 T 指针作为方法的结构体, 并没有实现 I 接口. TODO: 这里到底 实现了接口吗?
func (t *T) M() {
	fmt.Println(t.S)
}

// 定义别名类型 F
type F float64

// F 类型实现了 I 接口
func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I

	// (&{Hello}, *main.T)
	// Hello
	i = &T{"Hello"}
	printInfo(i)
	i.M()

	// (3.141592653589793, main.F)
	// 3.141592653589793
	i = F(math.Pi)
	printInfo(i)
	i.M()
}

func printInfo(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
