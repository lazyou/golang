package main

import "fmt"
import "time"

// TODO: 更费解了

// 有时候会出现goroutine阻塞的情况，那么我们如何避免整个程序进入阻塞的情况呢？
// 我们可以利用select来设置超时
func main() {
	c := make(chan int)
	o := make(chan bool)
	
	go func() {
		for {
			select {
				case v := <- c:
					fmt.Println(v)
				case <- time.After(2 * time.Second):
					fmt.Println("timeout")
					o <- true
					break
			}
		}
	}()

	result := <-o

	fmt.Println(result)
}