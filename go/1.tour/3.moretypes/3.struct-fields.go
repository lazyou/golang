package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main()  {
	v := Vertex{1, 2}

	// 结构体字段使用 **点号** 来访问
	v.X = 4
	fmt.Println(v.X)
}