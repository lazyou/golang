package main

import "fmt"

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func main() {
	sq1 := new(Square) // 使用实现接口的结构体
	sq1.side = 5
	fmt.Printf("The square has area: %f\n", sq1.Area())

	var areaIntf Shaper // 使用接口类型
	areaIntf = sq1
	fmt.Printf("The square has area: %f\n", areaIntf.Area())
}
