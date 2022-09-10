// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 33.
//!+

// Echo4 prints its command-line arguments.
package main

import (
	"flag"
	"fmt"
	"strings"
)

//参数1：命令行传递参数的名称  参数2：默认值   参数3：参数的说明
var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

// flag 包用来解析命令行参数
// eg: go run 2.3.echo4/main.go -s -- 1 2 3 // `-s --`设置 分割符号为 '--'
// go run 2.3.echo4/main.go -n 1 2 3 // `-n` 不忽略换行符， 无 `-n` 默认为 false，有则 true
func main() {
	flag.Parse()

	// flag.Args() 收到的参数， slice
	fmt.Print(strings.Join(flag.Args(), *sep))

	if !*n {
		fmt.Println()
	}
}

//!-
