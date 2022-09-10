package main

import (
	"fmt"
)

type Item interface {
	Name() string
	Price() float64
}

// VegBurger 结构体, 并实现了 Item 接口
type VegBurger struct {
}

func (r *VegBurger) Name() string {
	return "vegburger"
}

func (r *VegBurger) Price() float64 {
	return 1.5
}

// ChickenBurger 结构体, 并实现了 Item 接口
type ChickenBurger struct {
}

func (r *ChickenBurger) Name() string {
	return "chickenburger"
}

func (r *ChickenBurger) Price() float64 {
	return 5.5
}

func main() {
	var vb VegBurger
	fmt.Println(vb.Price())

	var cb ChickenBurger
	fmt.Println(vb.Price())

	useDuotai(&vb)
	useDuotai(&cb)
}

// 函数接收一个 接口作为参数, 所以只要实现了这个接口的类型都能被被传递进来
func useDuotai(i Item) {
	fmt.Printf("%T %v\n", i, i)
}
