package main

import (
	"fmt"
)

// 返回值可以是多个
func swap(x, y string) (string, string) {
    return y, x
}

func main()  {
    a, b := swap("hello", "world")

    fmt.Println(a, b)
}
