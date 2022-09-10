package main

import (
	"os"
	"fmt"
	"bufio"
)

// 从文件种读取内容, 并统计重复行
// `go run 1.3.2.dup2.go README.md`
func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		// 支持多个文件
		for _, arg := range files {
			f, err := os.Open(arg)

			// 文件打开失败, 例如文件不存在等情况
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}

			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		fmt.Printf("%d\t%s\n", n, line)
	}
}

func countLines(f *os.File, counts map[string]int)  {
	// 从文件中读取内容
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}