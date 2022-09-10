package main

import "fmt"

// 4.2.4append.go 优化版, 支持接收 slice 作为 append 参数
func main() {
	var x []int

	x = appendSlice(x, 1)
	x = appendSlice(x, 2, 3)
	x = appendSlice(x, 4, 5, 6)
	fmt.Println(x) // "[1 2 3 4 5 6 1 2 3 4 5 6]"
}

// “...” 省略号表示接收变长的参数为slice
func appendSlice(x []int, y ...int) []int {
	xlen := len(x)
	xcap := cap(x)
	ylen := len(y)

	var z []int
	zlen := xlen + ylen

	// 容量足够
	if zlen <= xcap {
		z = x[:zlen]
	} else {
		zcap := zlen

		if zcap < (2 * xlen) {
			zcap = 2 * xlen
		}

		z = make([]int, zlen, zcap)
		copy(z, x) // 前半部分赋值
	}

	copy(z[xlen:], y) // 后半部分赋值

	return z
}
