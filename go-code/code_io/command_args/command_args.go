package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	who := "Alice "

	// 从命令行读取参数
	if len(os.Args) >  1 {
		who += strings.Join(os.Args[1:], " ")
	}

	fmt.Println(who)
}