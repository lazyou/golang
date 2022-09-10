package main

import "fmt"

// 使用 range 迭代字符串时，需要注意的是 range 迭代的是 Unicode 而不是字节
// 第一个是被迭代的字符的 UTF-8 编码的第一个字节在字符串中的索引
// 第二个值的为对应的字符且类型为 rune(实际就是表示 unicode 值的整形数据）
func main() {
	const s = "Go语言"

	for i, r := range s {
		fmt.Printf("%#U : %d \n", r, i)
	}

	/*
	U+0047 'G' ： 0
	U+006F 'o' ： 1
	U+8BED '语' ： 2
	U+8A00 '言' ： 5
	 */

	// Unicode 字符
	var unicodeChar int = '\u8BED'
	fmt.Printf("%#U \n", unicodeChar) // U+8BED '语'
}
