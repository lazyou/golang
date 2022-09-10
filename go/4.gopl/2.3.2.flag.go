package main

import (
	"flag"
	"fmt"
	"strings"
)

// flag 包实现 命令行标记解析
// 调用flag.Bool函数会创建一个新的对应布尔型标志参数的变量:
// 第一个是的命令行标志参数的名字“n”，然后是该标志参数的默认值（这里是false），最后是该标志参数对应的描述信息。
var n = flag.Bool("n", false, "omit trailing newline -- 打印换行符号")
var sep = flag.String("s", " ", "separator -- 设置分隔符号")

func main() {
	// 当程序运行时，必须在使用标志参数对应的变量之前先调用flag.Parse函数，用于更新每个标志参数对应变量的值（之前是默认值）
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))

	// 是否省换行
	if !*n {
		fmt.Println()
	}

	/**
	$ go run 2.3.2.flag.go a bc def
	a bc def

	$ go run 2.3.2.flag.go -s / a bc def
	a/bc/def

	$ go run 2.3.2.flag.go -n a bc def
	a bc def$

	$ go run 2.3.2.flag.go -help
	Usage of 2.3.2.flag.go:
	  -n    omit trailing newline
	  -s string
			separator (default " ")
	 */
}