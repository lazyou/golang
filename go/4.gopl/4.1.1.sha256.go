package main

import (
	"crypto/sha256"
	"fmt"
)

// crypto/sha256 包的 Sum256 函数对一个任意的字节 slice 类型的数据生成一个对应的消息摘要。
func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	//fmt.Printf("%v \n", c1)

	// %x 副词参数，它用于指定以十六进制的格式打印数组或slice全部的元素，
	fmt.Printf("%x \n%x \n", c1, c2)

	// %t 副词参数是用于打印布尔型数据，
	fmt.Printf("%t \n", c1 == c2)

	// %T 副词参数是用于显示一个值对应的数据类型
	fmt.Printf("%T \n", c1)
}
