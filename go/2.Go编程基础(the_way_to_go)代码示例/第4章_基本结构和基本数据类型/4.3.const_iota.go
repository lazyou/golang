package main

import (
	"fmt"
)

// 常量用作枚举。 但是不推荐， 万一往中间插了一个常量后续的值都变了
const (
	Sunday  = iota // 0
	Monday         // 1
	Tuesday        // 2
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	fmt.Println(Sunday)
	fmt.Println(Monday)
	fmt.Println(Tuesday)
}
