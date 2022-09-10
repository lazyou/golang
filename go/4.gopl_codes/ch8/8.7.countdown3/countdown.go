// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 246.

// Countdown implements the countdown for a rocket launch.
package main

// NOTE: the ticker goroutine never terminates if the launch is aborted.
// This is a "goroutine leak".

import (
	"fmt"
	"os"
	"time"
)

//!+

func main() {
	// ...create abort channel...

	//!-

	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- struct{}{}
	}()

	//!+
	fmt.Println("Commencing countdown.  Press return to abort.")
	// time.Tick函数表现得好像它创建了一个在循环中调用time.Sleep的goroutine，每次被唤醒时发送一个事件
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case t := <-tick:
			fmt.Println(t)
			fmt.Println("tick!")
		case ab := <-abort:
			fmt.Println(ab) // 空结构体 `{}`
			fmt.Println("Launch aborted!")
			return
		}
	}
	launch()
}

//!-

func launch() {
	fmt.Println("Lift off!")
}
