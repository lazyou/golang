// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 229.

// Pipeline2 demonstrates a finite 3-stage pipeline.
package main

import (
	"fmt"
	"time"
)

//!+
func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		fmt.Println("Counter sleep...")
		time.Sleep(time.Duration(1)*time.Second)

		for x := 0; x < 10; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	// Squarer
	go func() {
		fmt.Println("Squarer sleep...")
		time.Sleep(time.Duration(1)*time.Second)

		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	// Printer (in main goroutine)
	fmt.Println("Printer sleep...")
	time.Sleep(time.Duration(1)*time.Second)

	for x := range squares {
		fmt.Println(x)
	}
}

//!-
