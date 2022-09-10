package main

import (
	"fmt"
)

// 2.2 数组 array : len() cap()
func main() {
	var arr [10]int; // 声明一个长度为10的int类型数组
	arr[0] = 10;
	arr[1] = 1;

	fmt.Printf("arr is %v\n", arr)
	fmt.Printf("first is %d\n", arr[0])
	fmt.Printf("last is %v\n", arr[9])

	a := [3]int{1, 2, 3}
	fmt.Printf("%v \n", a)

	c := [...]int{4, 5, 6}
	fmt.Printf("%v\n", c)
	fmt.Println(len(c))
	fmt.Println(cap(c))
}
