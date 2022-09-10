// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.

// Echo3 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strings"
)

//!+
func main() {
	// 使用 strings.Join 拼接字符串效率高，不会 "每次循环迭代字符串s的内容都会更新"
	fmt.Println(strings.Join(os.Args[1:], " "))
}

//!-
