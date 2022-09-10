package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func (p Person) String() string {
	return fmt.Sprintf("%f (%v years)", p.Name, p.Age)
}

/*
Stringers
一个普遍存在的接口是 `fmt` 包中定义的 Stringer。

type Stringer interface {
    String() string
}

Stringer 是一个可以用字符串描述自己的类型。`fmt`包 （还有许多其他包）使用这个来进行输出。
*/
func main()  {
	a := Person{"A D", 42}
	z := Person{"Z B", 9001}

	// %!f(string=A D) (42 years) %!f(string=Z B) (9001 years)
	fmt.Println(a, z)
}