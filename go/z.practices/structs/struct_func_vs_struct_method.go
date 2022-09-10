package main

import "fmt"

// 比较结构体指针作为函数参数 和 结构体指针作为方法的接收者 之间的区别

type Vertex struct {
	X, Y float64
}

// Vertex 结构体指针作为方法的接收者
func (v *Vertex) ScaleMethod(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Vertex 结构体指针作为函数参数
func ScaleFunction(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.ScaleMethod(2) // 结构体类型 调用方法 (理论上应该报错)
	ScaleFunction(&v, 10)

	p := &Vertex{3, 4}
	p.ScaleMethod(2) // 指针类型 调用方法
	ScaleFunction(p, 10)

	// {60 80} &{60 80}
	fmt.Println(v, p)

	// TODO: 更多笔记 -- 方法与指针重定向
	/*
		// 比较前两个程序，你大概会注意到带指针参数的函数 __必须接受一个指针__：
		var v Vertex
		ScaleFunc(v, 5)  // 编译错误！
		ScaleFunc(&v, 5) // OK


		// 而以指针为接收者的方法被调用时，接收者 __既能为值又能为指针__：
		var v Vertex
		v.Scale(5)  // OK -- Go 会将语句 v.Scale(5) 解释为 (&v).Scale(5)
		p := &v
		p.Scale(10) // OK

		// 对于语句 v.Scale(5)，即便 v 是个值而非指针，带指针接收者的方法也能被直接调用。 也就是说，由于 Scale 方法有一个指针接收者，为方便起见，Go 会将语句 v.Scale(5) 解释为 (&v).Scale(5)
	*/
}
