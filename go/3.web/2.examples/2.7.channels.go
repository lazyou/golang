package main

import "fmt"

func sum(a []int, c chan int) {
	total := 0
	
	for _, v := range a {
		total += v
	}

	c <- total // send total to c
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}
	aHalfLen := len(a) / 2

	// 必须使用make 创建channel
	c := make(chan int)

	fmt.Println(a[:aHalfLen])
	fmt.Println(a[aHalfLen:])

	go sum(a[:aHalfLen], c)
	go sum(a[aHalfLen:], c)

	// channel接收和发送数据都是阻塞的
	// 所谓阻塞，也就是如果读取（value := <-ch）它将会被阻塞，直到有数据接收
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x + y)
}