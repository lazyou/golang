package main

import "fmt"

// struct 的四种声明使用方式
func main() {
	type Person struct {
		name string
		age  int
	}

	// 1. 点符号赋值初始化
	var p1 Person
	p1.name = "zhang"
	p1.age = 19
	fmt.Println(p1)

	// 2. 按照顺序提供初始化值
	p2 := Person{
		"Tom",
		14,
	}
	fmt.Println(p2)

	// 3. 通过field:value的方式初始化，这样可以任意顺序
	p3 := Person{
		age:  33,
		name: "Any",
	}
	fmt.Println(p3)

	// 4. 通过 new函数 分配一个指针，此处 P 的类型为 `*person`
	p4 := new(Person)
	p4.name = "ppp4"
	p4.age = 73
	fmt.Println(p4)  // &{ppp4 73} 注意: 只有这边是指针类型
	fmt.Println(*p4) // {ppp4 73}
}
