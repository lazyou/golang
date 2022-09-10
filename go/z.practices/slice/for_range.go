package main

import (
	"fmt"
)

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	// for 循环的 range 形式可遍历切片或映射
	for index, value := range pow {
		fmt.Printf("index = %d; value = %d; \n", index, value)
	}
}
