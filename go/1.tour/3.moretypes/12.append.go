package main

import "fmt"

/*
向 slice 添加元素

向 slice 的末尾添加元素是一种常见的操作，因此 Go 提供了一个内建函数 append: func append(s []T, vs ...T) []T

append 的第一个参数 s 是一个元素类型为 T 的 slice ，其余类型为 T 的值将会附加到该 slice 的末尾。

append 的结果是一个包含原 slice 所有元素加上新添加的元素的 slice。
*/
func main()  {
	var a []int
	printSlice("a", a)

	a = append(a, 0)
	printSlice("a", a)
	
	a = append(a, 1)
	printSlice("a", a)

	a = append(a, 2, 3, 4)
	printSlice("a", a)
}

/*
a len=0 cap=0 []
a len=1 cap=1 [0]
a len=2 cap=2 [0 1]
a len=5 cap=6 [0 1 2 3 4]
*/
func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}