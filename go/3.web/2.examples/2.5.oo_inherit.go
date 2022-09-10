package main

import "fmt"

// 继承的实现 (本质是组合嵌套)
type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human  // 匿名字段 (继承)
	school string
}

type Employee struct {
	Human
	company string
}

// 这里 (h Human) 或者 (h *Human) 都是一样的 (所以没事别乱用指针类型, 看着还是有点蛋疼)
func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

// 重写: Employee 的 method 重写 Human 的 method
func (e Employee) SayHi(rewrite string) {
	fmt.Println(rewrite)

	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
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
	sam.SayHi("重写参数可以不一致")
}
