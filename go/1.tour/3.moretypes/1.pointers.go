package main

import "fmt"

/*
指针
	* Go 具有指针。 指针保存了变量的内存地址。

	* 类型 `*T` 是指向类型 T 的值的指针。其零值是 `nil` 。

	* `&` 符号会生成一个指向其作用对象的指针 (取地址符号?)

	* `*` 符号表示指针指向的底层的值
*/
func main()  {
	i, j := 42, 2701

	p := &i
	fmt.Println(*p) // 42

	*p = 21
	fmt.Println(i)  // 21, i的值通过指针所指向的地址被改变了

	p = &j
	*p = *p / 37
	fmt.Println(j) // 2701 / 37
}
