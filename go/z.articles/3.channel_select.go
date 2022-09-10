package main

import (
	"fmt"
	//"time"
)

func send(c chan int) {
	for i := 1; i < 10; i++ {
		c <- i

		fmt.Println("send data: ", i)
	}
}

func main() {
	resch := make(chan int, 20)
	strch := make(chan string, 20)

	go send(resch)

	strch <- "wd"
	//time.Sleep(time.Second * 1)

	select {
	case a := <-resch:
		fmt.Println("get data:", a)
	case b := <-strch:
		fmt.Println("get data:", b)
	default:
		fmt.Println("no channel actvie")
	}
}
