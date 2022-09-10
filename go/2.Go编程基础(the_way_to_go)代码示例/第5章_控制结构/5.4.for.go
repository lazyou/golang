package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Printf("This is the %d iteration\n", i)
	}

	// 使用 for 循环迭代一个 Unicode 编码的字符串
	// ASCII 编码的字符占用 1 个字节，既每个索引都指向不同的字符，而非 ASCII 编码的字符（占有 2 到 4 个字节）不能单纯地使用索引来判断是否为同一个字符。我们会在第 5.4.4 节解决这个问题
	str2 := "日本語"
	fmt.Printf("The length of str2 is: %d\n", len(str2))
	for ix := 0; ix < len(str2); ix++ {
		fmt.Printf("Character on position %d is: %c \n", ix, str2[ix])
	}

	// while
	var i int = 5
	for i >= 0 {
		i = i - 1
		fmt.Printf("The variable i is now: %d\n", i)
	}
}
