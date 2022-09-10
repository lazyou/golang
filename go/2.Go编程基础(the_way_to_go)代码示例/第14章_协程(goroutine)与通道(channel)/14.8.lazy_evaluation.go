package main

import (
	"fmt"
)

var resume chan int

// 惰性生成器的实现
func main() {
	resume = integers()

	fmt.Println(generateInteger()) //=> 0
	fmt.Println(generateInteger()) //=> 1
	fmt.Println(generateInteger()) //=> 2
}

func integers() chan int {
	yield := make(chan int)
	count := 0

	go func() {
		for {
			yield <- count
			count++
		}
	}()

	return yield
}

func generateInteger() int {
	return <-resume
}
