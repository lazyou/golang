package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	// 没有条件的 switch: 没有条件的 switch 同 switch true 一样
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
