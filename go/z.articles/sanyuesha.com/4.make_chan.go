package main

import "fmt"

func main() {

	var c1 chan string
	// chan 零值是 nil
	if c1 == nil {
		fmt.Printf("c1 is nil --> %#v \n ", c1) //(chan string)(nil)
	}

	c2 := make(chan string)
	if c2 == nil {
		fmt.Printf("c2 is nil --> %#v \n ", c2)
	} else {
		// make 产生的 chan 不再是零值
		fmt.Printf("c2 is not nill --> %#v \n ", c2) //(chan string)(0xc42007c060)
	}
}
