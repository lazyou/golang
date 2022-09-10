// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.
//!+

// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)

	// os.Stdin 标准输入
	input := bufio.NewScanner(os.Stdin)

	// 每次调用 input.Scan()，即读入下一行，并移除行末的换行符
	for input.Scan() {
		// 读取的内容可以调用 input.Text() 得到
		counts[input.Text()]++
	}

	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

	// fmt.Printf("%T\n %v\n", input, input)
}

//!-
