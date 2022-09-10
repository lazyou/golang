// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 12.

//!+

// Dup3 prints the count and text of lines that
// appear more than once in the named input files.
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)

	for _, filename := range os.Args[1:] {
		// ReadFile 函数返回一个字节切片（byte slice），必须把它转换为string，才能用 strings.Split 分割。
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}

		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	fmt.Printf("重复次数\t重复内容\n")
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t\t%s\n", n, line)
		}
	}
}

//!-
