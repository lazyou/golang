package main

import "fmt"

func main() {
	mychannel := make(chan int, 10)

	for i := 0; i < 10; i++ {
		mychannel <- i
	}

	close(mychannel) // 关闭管道

	fmt.Println("data length: ", len(mychannel)) // 10

	for v := range mychannel {
		fmt.Println(v)
	}

	fmt.Printf("data lenght:  %d", len(mychannel)) // 0
}
