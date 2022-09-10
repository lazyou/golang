package main

import (
	"fmt"
)

func main()  {
    // 返回地址(也就是指针)
    p1 := new(int)
    fmt.Println(p1)     // 0xc420012090: 指针, 内存地址
    fmt.Println(*p1)    // 0, 指针指向的值, 默认零值

    var p2 *int;
    i := 0
    p2 = &ii
    fmt.Println(p2)
    fmt.Println(*p2)
}
