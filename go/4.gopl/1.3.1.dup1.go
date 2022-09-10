package main

import (
	"bufio"
	"os"
	"fmt"
)

// 重复的行统计
// 终端情况下请使用 ctrl+d 退出程序(输入界面)
func main() {
	counts := make(map[string]int)
	// Scanner类型是该包最有用的特性之一，它读取输入并将其拆成行或单词
	input := bufio.NewScanner(os.Stdin)

	// 每次调用input.Scan()，即读入下一行，并移除行末的换行符；
	// 读取的内容可以调用input.Text()得到。Scan函数在读到一行时返回true，不再有输入时返回false
	for input.Scan() {
		// 一下两行可以简化成: `counts[input.Text()]++`
		line := input.Text()
		counts[line] = counts[line] + 1 // key 是行内容, value 是行重复次数

		fmt.Println("用户输入的行内容:", line)
	}

	for line, n:= range counts{
		fmt.Printf("行内容 %s 重复 %d 次\n", line, n)
	}
}