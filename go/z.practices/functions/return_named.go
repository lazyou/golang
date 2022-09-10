package main

import (
	"fmt"
)

// 命名返回值
func swap(x, y string) (a, b string) {
    a = y
    b = x
    return 
}

func main()  {
    a, b := swap("hello", "world")

    fmt.Println(a, b)
}
