package main

import (
	"fmt"
)

func main() {
	// 类型 [n]T 表示拥有 n 个 T 类型的值的数组
	var str [2]string

	str[0] = "Hello"
	str[1] = "World"

	// [Hello World]
	fmt.Println(str)
	// Hello World
	fmt.Println(str[0], str[1])
	// Type: [2]string
	fmt.Printf("Type: %T Value: %v\n", str, str)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	// Type: [6]int
	fmt.Printf("Type: %T Value: %v\n", primes, primes)
}
