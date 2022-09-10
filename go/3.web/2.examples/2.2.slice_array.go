package main

import (
	"fmt"
)

func main() {
	/*
	 * array 和 slice
	 */
	// array: 10 个元素的 byte 数组
	var ar = [10]string {"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	
	// slice
	var a, b []string

	a = ar[2:5] // slice 指向数组的第 2 ~ 第 5 个元素 

	b = ar[3:5]

	// [a b c d e f g h i j]
	// [c d e]
	// [d e]
	fmt.Println(ar)
	fmt.Println(a)
	fmt.Println(b)

	// slice是引用类型，所以当引用改变其中元素的值时，其它的所有引用都会改变该值
	ar[4] = "eee"
	// [a b c d eee f g h i j]
	// [c d eee]
	// [d eee]
	fmt.Println(ar)
	fmt.Println(a)
	fmt.Println(b)
}