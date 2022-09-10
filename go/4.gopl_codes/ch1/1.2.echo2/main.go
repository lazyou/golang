// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 6.
//!+

// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""

	// for循环的另一种形式, 在某种数据类型的区间（range）上遍历，如字符串或切片
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

//!-
