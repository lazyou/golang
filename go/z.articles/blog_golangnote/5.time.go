package main

import (
	"fmt"
	"time"
)

func main()  {
    // 获取当前(当地)时间
    t := time.Now()
    fmt.Println(t)

    // 获取0时区时间
    t = time.Now().UTC()
    fmt.Println(t)

    // 获取当前时间戳
    timestamp := t.Unix()
    fmt.Println(timestamp)

    // 获取时区信息
    name, offset := t.Zone()
    fmt.Println(name, offset)

    // 格式化时间: TODO:格式化没有像 Y-m-d H:i:s 的做法吗?
    currenttime := time.Unix(timestamp + int64(offset), 0)
    fmt.Println("Current time:", currenttime.Format("2006-01-02 15:04:05"))
}
