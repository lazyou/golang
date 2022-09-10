package main

import "fmt"

// Vertex 坐标
type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}

	// 1: 结构体指针
	p := &v

	// TODO: 这里明明只是取地址, 怎么就能直接赋值了, 应该是 *p.X = 99 来赋值才对
	// __不过这么写太啰嗦了，所以语言也允许我们使用 __隐式间接引用__， 直接写 p.X 就可以__
	p.X = 99

	fmt.Println(v)
	// Type: *main.Vertex Value: &{99 2}
	fmt.Printf("Type: %T Value: %v\n", p, p)

	// 2: 结构体指针
	VPointer1 := &Vertex{}
	VPointer2 := &Vertex{}

	VPointer1.X = 1
	VPointer1.Y = 1

	// Type: *main.Vertex Value: &{1 1}
	fmt.Printf("Type: %T Value: %v\n", VPointer1, VPointer1)

	// Type: *main.Vertex Value: &{0 0}
	fmt.Printf("Type: %T Value: %v\n", VPointer2, VPointer2)
}
