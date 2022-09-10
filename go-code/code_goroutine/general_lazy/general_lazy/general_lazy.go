package main

import "fmt"

var resume chan int

func integers() chan int {
	yield := make(chan int)
	count := 0

	// TODO: 这样疯狂自增不浪费资源码? 应该是会的
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

func main() {
	resume = integers()
	fmt.Println(generateInteger())  //=> 0
	fmt.Println(generateInteger())  //=> 1
	fmt.Println(generateInteger())  //=> 2
}