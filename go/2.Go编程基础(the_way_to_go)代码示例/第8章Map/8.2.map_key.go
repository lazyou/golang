package main

import "fmt"

func main() {
	map1 := make(map[string]int)
	_, ok := map1["key1"] // 如果key1存在则ok == true，否则ok为false
	if ok {
		fmt.Println("key1 存在")
	} else {
		fmt.Println("key1 不存在")
	}
}
