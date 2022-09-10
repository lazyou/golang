package main

import (
	"fmt"
	"time"
)

const LIM = 41

var fibs [LIM]uint64
var count int = 0

func main() {
	var result uint64 = 0
	start := time.Now()
	for i := 0; i < LIM; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}

// 通过在内存中缓存和重复利用相同计算的结果，称之为内存缓存
func fibonacci(n int) (res uint64) {
	// 已经计算过的情况放在内存
	// memoization: check if fibonacci(n) is already known in array:
	if fibs[n] != 0 {
		res = fibs[n]
		return
	}

	count++
	fmt.Printf("fibonacci 被调用第 %d 次\n", count)

	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}

	fibs[n] = res

	return
}
