package main

import "fmt"

/*
构造 slice
	slice 由函数 `make` 创建。这会分配一个全是零值的数组并且返回一个 slice 指向这个数组
		a := make([]int, 5)  // len(a)=5

	为了指定**容量**，可传递第三个参数到 make
		b := make([]int, 0, 5) // len(b)=0, cap(b)=5
*/
func main()  {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5) // 长度0, 容量5 这就很费解了!
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	// cap()函数返回的是数组切片分配的空间大小
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
