package main

import "fmt"

func main() {
	sum := 0

	// Go **只有一种** 循环结构—— for 循环
	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println(sum)
}
