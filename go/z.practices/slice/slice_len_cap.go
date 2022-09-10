package main

import (
	"fmt"
)

func main() {
	// 切片的 __长度__ 就是它所包含的元素个数
	// 切片的 __容量__ 是从它的第一个元素开始数，到其底层数组元素末尾的个数
	s := []int{2, 3, 5, 7, 11, 13}
	// len=6 cap=6 [2 3 5 7 11 13]
	printSlice(s)

	s = s[:0]
	// len=0 cap=6 []
	printSlice(s)

	s = s[:4]
	// len=4 cap=6 [2 3 5 7]
	printSlice(s)

	s = s[2:]
	// len=2 cap=4 [5 7]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
