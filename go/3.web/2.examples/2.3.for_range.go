package main

import (
	"fmt"
)

func main()  {
	map1 := map[string]int{
		"one": 1,
		"two": 2,
	}

	// 遍历 map: key, val
	for k, v := range map1 {
		fmt.Println("key:", k)
		fmt.Println("val:", v)
	}

	// 省略 key
	for _, v := range map1 {
		fmt.Println("val:", v)
	}
}