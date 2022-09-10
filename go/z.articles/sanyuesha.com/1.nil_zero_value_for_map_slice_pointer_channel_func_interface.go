package main

import "fmt"

func main() {
	var m map[int]string
	var ptr *int
	var c chan int
	var sl []int
	var f func()
	var i interface{}

	// map[int]string(nil)
	fmt.Printf("%#v\n", m)

	// (*int)(nil)
	fmt.Printf("%#v\n", ptr)

	// (chan int)(nil)
	fmt.Printf("%#v\n", c)

	// []int(nil)
	fmt.Printf("%#v\n", sl)

	// (func())(nil)
	fmt.Printf("%#v\n", f)

	// <nil>
	fmt.Printf("%#v\n", i)
}
