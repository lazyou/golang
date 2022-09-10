package main

import (
	"unicode/utf8"
	"bufio"
	"os"
	"io"
	"fmt"
	"unicode"
)

// 字符统计
func main() {
	counts := make(map[rune]int)
	// UTF-8编码的长度总是从1到utf8.UTFMax（最大是4个字节）
	var utflen [utf8.UTFMax + 1]int
	invalid := 0

	in := bufio.NewReader(os.Stdin)

	for {
		// ReadRune方法执行UTF-8解码并返回三个值：解码的rune字符的值，字符UTF-8编码后的长度，和一个错误值。
		r, n, err := in.ReadRune()

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}

		// 如果输入的是无效的UTF-8编码的字符，返回的将是unicode.ReplacementChar表示无效字符，并且编码长度是1
		if r == unicode.ReplacementChar && n == 1 {
			invalid ++
			continue
		}

		counts[r]++
		utflen[n]++
	}

	fmt.Printf("rune \t count \n")
	for c, n := range counts {
		fmt.Printf("%q \t %d \n", c, n)
	}

	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}

	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
