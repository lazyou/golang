package main

import (
	"fmt"
)

// 返回值的类型写在方法名后面
func add(x int, y int) int {
    return x + y
}

func main()  {
    fmt.Println(add(42, 13))
}
