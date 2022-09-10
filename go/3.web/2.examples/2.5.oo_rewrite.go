package main

import "fmt"

// 重写
type Human struct {
	name  string
	age   int
	phone string
}

// Human 定义 method
func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

type Student struct {
	Human  //匿名字段
	school string
}

type Employee struct {
	Human   //匿名字段
	company string
}

// Employee 的 method __重写__ Human 的 method
func (e *Employee) SayHi() {
	fmt.Printf("Employee 重写了 Human SayHi 方法 -- Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone) //Yes you can split into 2 lines here.
}

func main() {
	mark := Student{
		Human{
			"Mark",
			25,
			"222-222-YYYY",
		},
		"MIT",
	}

	sam := Employee{
		Human{
			"Sam",
			45,
			"111-888-XXXX",
		},
		"Golang Inc",
	}

	mark.SayHi()
	sam.SayHi()
}
