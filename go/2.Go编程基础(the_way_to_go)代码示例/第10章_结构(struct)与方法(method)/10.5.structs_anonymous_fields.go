package main

import "fmt"

type innerS struct {
	in1 int
	in2 int
}

type outerS struct {
	b      int
	c      float32
	int    // anonymous field
	innerS //anonymous field
}

// 匿名字段
func main() {
	outer := new(outerS)
	outer.b = 6
	outer.c = 7.5
	outer.int = 60 // 匿名字段
	outer.in1 = 5  // 继承字段,且匿名字段
	outer.in2 = 10 // 继承字段,且匿名字段

	fmt.Printf("outer.b is: %d\n", outer.b)
	fmt.Printf("outer.c is: %f\n", outer.c)
	fmt.Printf("outer.int is: %d\n", outer.int)
	fmt.Printf("outer.in1 is: %d\n", outer.in1)
	fmt.Printf("outer.in2 is: %d\n", outer.in2)

	// 使用结构体字面量
	outer2 := outerS{6, 7.5, 60, innerS{5, 10}}
	fmt.Println("outer2 is:", outer2)
}
