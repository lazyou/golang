package main

import (
	"fmt"
	"time"
)

func addele(a []int, i int) {
	defer func() { // 匿名函数捕获错误
		err := recover()

		if err != nil {
			//fmt.Println(err)
			fmt.Println("add ele fail")
		}
	}()

	a[i] = i
	fmt.Println(a)
}

func main() {
	arr := make([]int, 4)

	for i := 0; i < 10; i++ {
		go addele(arr, i)
	}

	fmt.Printf("final arr1 : %v \n", arr) // 与 goroutine 并行 (异步执行), arr 每个字段的值不一定

	time.Sleep(time.Second * 1)

	fmt.Printf("final arr2 : %v \n", arr) // sleep 阻塞之后, 所以一定是 final arr2 : [0 1 2 3]
}
