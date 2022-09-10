package main
import (
	"fmt"
	"time"
)

func main() {
	ch := pump2()
	suck2(ch)

	time.Sleep(1e9)
}

func pump2() chan int {
	ch := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
		}
	}()

	return ch
}

func suck2(ch chan int) {
	go func() {
		// 给通道使用 for 循环
		for v := range ch {
			fmt.Println(v)
		}
	}()
}
