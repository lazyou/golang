package main

import "fmt"

type Foo struct {
	name string
	age  int
}

// 以下代码演示了 struct 初始化的过程，可以说明不使用 new 一样可以完成 struct 的初始化工作。
func main() {
	// 声明初始化
	var foo1 Foo
	fmt.Printf("foo1 --> %#v\n ", foo1) //main.Foo{age:0, name:""}
	foo1.age = 1
	fmt.Println(foo1.age)

	//struct literal 初始化
	foo2 := Foo{}
	fmt.Printf("foo2 --> %#v\n ", foo2) //main.Foo{age:0, name:""}
	foo2.age = 2
	fmt.Println(foo2.age)

	// 指针初始化
	foo3 := &Foo{}
	fmt.Printf("foo3 --> %#v\n ", foo3) //&main.Foo{age:0, name:""}
	foo3.age = 3                        // foo3 是指针类型, 理论上应该 (*foo3).age = 3 才能赋值 -- TODO: 这里 go 自动做了转化
	fmt.Println(foo3.age)

	// new 初始化
	foo4 := new(Foo)
	fmt.Printf("foo4 --> %#v\n ", foo4) //&main.Foo{age:0, name:""}
	foo4.age = 4
	fmt.Println(foo4.age)

	//声明指针并用 new 初始化
	var foo5 *Foo = new(Foo)
	fmt.Printf("foo5 --> %#v\n ", foo5) //&main.Foo{age:0, name:""}
	foo5.age = 5
	fmt.Println(foo5.age)
}
