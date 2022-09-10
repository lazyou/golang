package main

import "fmt"

// 2.2 变量
func main()  {
	var a int
	var b, c int
	var d int = 4
	var e, f int = 5, 6
	var g, h = 7, 8 // Go会根据其相应值的类型来帮你初始化它们

	a = 1
	b = 2
	c = 3

	_, i := 0, 9

	fmt.Println(a)
	fmt.Println(b, c)
	fmt.Println(d)
	fmt.Println(e, f)
	fmt.Println(g, h)
	fmt.Println(i)
}
