package main

import "fmt"

// 实现一个 append, 专门用于处理[]int类型的slice, 用于理解slice底层是如何工作
func main() {
	var x, y []int

	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}

	/*
		0 cap=1 [0]
		1 cap=2 [0 1]
		2 cap=4 [0 1 2]
		3 cap=4 [0 1 2 3]
		4 cap=8 [0 1 2 3 4]
		5 cap=8 [0 1 2 3 4 5]
		6 cap=8 [0 1 2 3 4 5 6]
		7 cap=8 [0 1 2 3 4 5 6 7]
		8 cap=16        [0 1 2 3 4 5 6 7 8]
		9 cap=16        [0 1 2 3 4 5 6 7 8 9]
	*/
}

// 每次调用appendInt函数，必须先检测slice底层数组是否有足够的容量来保存新添加的元素
func appendInt(x []int, y int) []int {
	xlen := len(x)
	xcap := cap(x)

	var z []int
	zlen := xlen + 1

	// 容量足够
	if zlen <= xcap {
		z = x[:zlen]
	} else {
		// 容量计算, 设置为至少 xlen 的两倍
		zcap := zlen

		if zcap < (2 * xlen) {
			zcap = 2 * xlen
		}

		z = make([]int, zlen, zcap)
		copy(z, x) // 内置的copy函数可以方便地将一个slice复制另一个相同类型的slice
	}

	z[xlen] = y

	return z
}
