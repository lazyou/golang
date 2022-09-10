package main

import (
	"time"
	"fmt"
)

// TODO: 有时候会看到一行输出, 偶尔会看到两行输出(比较少)
func main()  {
    go fmt.Println("goroutine message")

    // 让新开启的 goroutine 有机会得到执行，开启一个 goroutine 之后，后续的代码会继续执行，
    // 后续代码执行完毕程序就终止了，而开启的 goroutine 可能还没开始执行 (这就是为什么 有时候会看到一行输出, 偶尔会看到两行输出(比较少))
    time.Sleep(1) 

    fmt.Println("main function message")
}
