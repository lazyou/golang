package main

import (
	"fmt"
)

// Vertex 坐标
type Vertex struct {
	// 字段没有大写也能被访问: 只不过 go 提倡首字母大写表示能被访问, 所以也应用到这里
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}

	// Type: main.Vertex Value: {1 2}
	fmt.Printf("Type: %T Value: %v\n", v, v)

	fmt.Println(v)

	fmt.Println(v.X)
}
