package main

import (
	"fmt"
)

type Person struct{
    name string
    age int
}

func (p *Person) sayName() {
    fmt.Println(p.name)    
}

// 有点像继承
type Student struct {
    Person
    name string
}

func main() {
    p := Person{name:"C"}
    p.sayName() // #1 C

    // Student 虽然可以直接调用 embed type 的 method，但是 method 的 receiver 仍然是 embed type，所以 #2 输出为空。
    s1 := Student{name:"Java"}
    s1.sayName() // #2 此行输出空字符串

    s2 := Student{
        name: "Java",
        Person: Person {
            name: "VB",
        },
    }
    s2.sayName()            // #3 VB

    // Student 中的同名属性可以遮蔽 embed type 的属性，#4 的输出
    fmt.Println(s2.name)    // #4 Java
    fmt.Println(s2.Person.name) //#5 VB
}
