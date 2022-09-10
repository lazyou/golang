package main

import "fmt"

// for-range 是 Go 特有的一种的迭代结构, 它可以迭代任何一个集合 (包括数组和 map)
func main() {
	str2 := "Chinese: 日本語"

	fmt.Printf("The length of str2 is: %d\n", len(str2))
	for pos, char := range str2 {
		fmt.Printf("character %c starts at byte position %d\n", char, pos)
	}

	fmt.Println()

	fmt.Println("index int(rune) rune    char bytes")
	for index, rune := range str2 {
		fmt.Printf("%-2d      %d      %U '%c' % X\n", index, rune, rune, rune, []byte(string(rune)))
	}
}
