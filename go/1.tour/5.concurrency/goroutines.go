package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	// TODO: 只执行这么一句不行吗? 没看到任何输出, 必须配合下面才能看到
	go say("go say")

	say("say")
}
