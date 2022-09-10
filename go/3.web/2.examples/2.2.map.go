package main

import (
	"fmt"
)

func main() {
	// 1. 声明 map
	var map1 map[string]int
	fmt.Println(map1)

	map1 = map[string]int{"one": 1} // 感觉很繁琐, 为撒子不能省略 `map[string]int`, 明明上面已经声明了?
	map1["two"] = 2
	fmt.Println(map1)

	// 2. map 的另一种声明方式
	map2 := make(map[string]int)
	// fmt.Println(map2)
	map2["one"] = 1
	map2["two"] = 2
	fmt.Println(map2)

	delete(map2, "one")
	fmt.Println(map2)

	/**
	使用map过程中需要注意的几点：
		map 是 __无序__ 的，每次打印出来的 map 都会不一样, 它不能通过index获取，而必须通过key获取;

		map 的长度是不固定的，也就是和 slice 一样，也是一种 __引用类型__;

		map 和其他基本型别不同，它不是 thread-safe，在多个 go-routine 存取时，必须使用 `mutex lock` 机制.
	 */

	// 初始化一个字典
	rating := map[string]float32{
		"C":      5,
		"Go":     4.5,
		"Python": 4.5,
		"C++":    2,
	}
	// map有两个返回值，第二个返回值，如果不存在key，那么ok为false，如果存在ok为true
	csharpRating, ok := rating["C#"]

	if ok {
		fmt.Println("C# is in the map and its rating is ", csharpRating)
	} else {
		fmt.Println("We have no rating associated with C# in the map")
	}

	delete(rating, "C")
}