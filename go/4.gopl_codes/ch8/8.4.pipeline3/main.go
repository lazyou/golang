// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 231.

// Pipeline3 demonstrates a finite 3-stage pipeline
// with range, close, and unidirectional channel types.
package main

import (
	"fmt"
	"time"
)

//!+
func counter(out chan<- int) {
	fmt.Println("counter sleep...")
	time.Sleep(time.Duration(1)*time.Second)

	for x := 0; x < 10; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	fmt.Println("squarer sleep...")
	time.Sleep(time.Duration(1)*time.Second)

	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int) {
	fmt.Println("printer sleep...")
	time.Sleep(time.Duration(1)*time.Second)

	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)

	/**
	printer sleep...
	counter sleep...
	squarer sleep...
	0
	1
	4
	9
	16
	25
	36
	49
	64
	81
	 */
}

//!-
