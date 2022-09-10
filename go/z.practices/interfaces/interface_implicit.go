package main

import (
	"fmt"
)

type I interface {
	M()
}

type T struct {
	S string
}

// 此方法表示类型 T 实现了接口 I，但我们无需显式声明此事
// 隐式接口从接口的实现中解耦了定义，这样接口的实现可以出现在任何包中，无需提前准备。
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	// 使用接口: 声明 i 是 I 接口类型, 然后从 I 接口的实现结构体 T 中获取
	var i I = T{"hello"}
	i.M()
}
