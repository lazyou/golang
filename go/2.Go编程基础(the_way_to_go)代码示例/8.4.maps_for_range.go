package main

import "fmt"

// 想获取一个 map 类型的切片，我们必须使用两次 make() 函数，第一次分配切片，第二次分配 切片中每个 map 元素
func main() {
	items := make([]map[int]int, 5)

	for i := range items {
		items[i] = make(map[int]int, 1)
		items[i][1] = 2
	}

	// Version A: Value of items: [map[1:2] map[1:2] map[1:2] map[1:2] map[1:2]]
	fmt.Printf("Version A: Value of items: %v\n", items)

	// Version B: NOT GOOD!
	items2 := make([]map[int]int, 5)
	for _, item := range items2 {
		item = make(map[int]int, 1) // item is only a copy of the slice element.
		item[1] = 2                 // This 'item' will be lost on the next iteration.
	}

	// Version B: Value of items: [map[] map[] map[] map[] map[]]
	fmt.Printf("Version B: Value of items: %v\n", items2)
}
