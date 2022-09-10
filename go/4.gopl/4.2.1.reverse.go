package main

import "fmt"

// 一个slice由三个部分构成：指针、长度和容量。指针指向第一个slice元素对应的底层数组元素的地址
func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}

	reverse(a[:])
	fmt.Println(a) // [5 4 3 2 1 0], a 被改变, 说明了 slice 是指针类型

	reverse(a[3:])
	fmt.Println(a) // [5 4 3 0 1 2], 只有部分被翻转, 说明了 slice 的底层确实引用一个数组对象
}

// 数组翻转
func reverse(s []int) {
	// 思路: i 从左边索引值递增, j 从右边索引值递减, 相互交换索引所在的值, 直到 i j 相遇 (也就是 i < j)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
