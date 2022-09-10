package main

import "fmt"

// 结构体: 一个结构体（ struct ）就是 **一个字段** 的集合
type Vertex struct {
	X int
	Y int
}

func main() {
	// {1 2}
	fmt.Println(Vertex{1, 2})
}