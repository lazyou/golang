package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// 切片: 切片并不存储任何数据，它只是描述了底层数组中的一段
	var s []int = primes[1:4]
	// Type: []int -- 显然和数组的差别就是 slice 没有限制数组的大小
	fmt.Printf("Type: %T Value: %v\n", s, s)

	primes[1] = 999
	fmt.Println(s) // 这里 slice 对应的 [1] 也被修改了 -- 证明了 切片就像数组的引用
}
