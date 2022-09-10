package main

import "fmt"

type TwoInts struct {
	a int
	b int
}

// Go 方法是作用在接收者（receiver）上的一个函数，接收者是某种类型的变量。 因此方法是一种特殊类型的函数。
// 最后接收者不能是一个指针类型，但是它可以是任何其他允许类型的指针。
// 类型 T（或 *T）上的所有方法的集合叫做类型 T（或 *T）的方法集。
// 如果 receiver 是一个指针，Go 会自动解引用
func main() {
	two1 := new(TwoInts)
	two1.a = 12
	two1.b = 10

	fmt.Printf("The sum is: %d\n", two1.AddThem())
	fmt.Printf("Add them to the param: %d\n", two1.AddToParam(20))

	two2 := TwoInts{3, 4}
	fmt.Printf("The sum is: %d\n", two2.AddThem())
}

func (tn TwoInts) AddThem() int {
	return tn.a + tn.b
}

func (tn *TwoInts) AddToParam(param int) int {
	return tn.a + tn.b + param
}
