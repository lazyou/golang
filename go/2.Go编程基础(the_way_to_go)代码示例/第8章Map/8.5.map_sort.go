// the telephone alphabet:
package main

import (
	"fmt"
	"sort"
)

var (
	barVal = map[string]int{
		"alpha": 34, "bravo": 56, "charlie": 23,
		"delta": 87, "echo": 56, "foxtrot": 12,
		"golf": 34, "hotel": 16, "indio": 87,
		"juliet": 65, "kili": 43, "lima": 98,
	}
)

func main() {
	fmt.Println("unsorted:")
	for k, v := range barVal {
		fmt.Printf("Key: %v, Value: %v / \n", k, v)
	}
	keys := make([]string, len(barVal))
	i := 0
	for k, _ := range barVal {
		keys[i] = k
		i++
	}

	// 字符串排序. 本质是提取map里面的key，对key排序后再循环key。
	// 并不是对 map 直接排序。
	sort.Strings(keys)
	fmt.Println()
	fmt.Println("sorted:")
	for _, k := range keys {
		fmt.Printf("Key: %v, Value: %v / \n", k, barVal[k])
	}
}
