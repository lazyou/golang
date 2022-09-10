package main

import "fmt"

// 内置的 append 函数用于向 slice 追加元素：
func main() {
	var runes []rune

	for _, r := range "Hello, 世界" {
		fmt.Println(r)        // 字符对应的 int
		fmt.Printf("%q\n", r) // 看到字符本身
		runes = append(runes, r)
	}

	fmt.Printf("%q\n", runes) // ['H' 'e' 'l' 'l' 'o' ',' ' ' '世' '界']
}
