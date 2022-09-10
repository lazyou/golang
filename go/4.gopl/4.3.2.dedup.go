package main

import (
	"bufio"
	"os"
	"fmt"
)

// 下面的dedup程序读取多行输入，但是只打印第一次出现的行
func main()  {
	seen := make(map[string]bool)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}

	fmt.Println(seen)

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}
}