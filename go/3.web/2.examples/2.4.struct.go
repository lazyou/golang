package main

import "fmt"

// 声明一个新的类型
type person struct {
	name string
	age  int
}

// 比较两个人的年龄，返回年龄大的那个人，并且返回年龄差
// struct也是传值的
func Older(p1 person, p2 person) (person, int) {
	if p1.age > p2.age { // 比较p1和p2这两个人的年龄
		return p1, p1.age - p2.age
	}

	return p2, p2.age - p1.age
}

func main() {
	var tom person

	// 赋值初始化
	tom.name, tom.age = "Tom", 18

	// 两个字段都写清楚的初始化
	bob := person{
		age:  26,
		name: "Bob",
	}

	// 按照struct定义顺序初始化值
	// bob := person{"Paul", 43}

	theOlder, thediff := Older(tom, bob)

	fmt.Printf("年长的是 %s, 大对方 %d 岁\n", theOlder.name, thediff)
}
