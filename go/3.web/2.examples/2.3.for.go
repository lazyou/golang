package main

import "fmt"

// 2.3 for 的两种用法
func main() {
	// 1.
	sum := 0
	for index := 0; index < 10; index++ {
		sum += index
	}
	fmt.Println("sum is equal to ", sum)

	// 2. while 用法
	sum2 := 1
	for sum2 < 1000 {
		sum2 += sum2
	}
	fmt.Println(sum2)
}
