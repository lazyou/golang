package main

import "fmt"

var count int = 0

func main() {
	result := 0
	for i := 0; i <= 3; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
}

func fibonacci(n int) (res int) {
	count++
	fmt.Printf("fibonacci 被调用第 %d 次\n", count) // 可以说是相当恐怖了

	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}

	return
}
