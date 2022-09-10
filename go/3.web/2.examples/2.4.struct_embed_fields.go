package main

import "fmt"

type Skills []string

type Human struct {
	name   string
	age    int
	weight int
}

// 说成 **继承** 更好理解
type Student struct {
	Human      // 匿名字段，那么默认Student就包含了Human的所有字段.
	Skills     // 匿名字段，自定义的类型string slice
	int        // 内置类型作为匿名字段 (TODO: 这个方式我看的费解, 尽量不这么用)
	speciality string
}

func main() {
	mark := Student{
		Human: Human{
			"Lxl",
			25,
			120,
		},
		speciality: "Computer Science",
	}

	// 我们访问相应的字段
	fmt.Println("His name is ", mark.name)
	fmt.Println("His age is ", mark.age)
	fmt.Println("His weight is ", mark.weight)
	fmt.Println("His speciality is ", mark.speciality)

	mark.name = "name 直接修改操作(不用通过 Human)"
	fmt.Println("His name is ", mark.name)

	mark.Human.age = -1
	fmt.Println("通过 Human 改 age, 但不是必须 ", mark.age)

	// 我们来修改他的skill技能字段
	mark.Skills = []string{"anatomy"}
	mark.Skills = append(mark.Skills, "physics", "golang")
	fmt.Println("Her skills are ", mark.Skills)

	// 修改匿名内置类型字段 (TODO: 就是看起来怪怪的, 像是有个到底是int类型, 还是 int 属性名)
	mark.int = 3
	fmt.Println("Her preferred number is", mark.int)
}
