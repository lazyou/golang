package main

import "fmt"

// TODO: slice map chan 是 __引用类型__
func main() {
	// slice
	s2 := make([]int, 3)
	fmt.Printf("%#v", s2) //[]int{0, 0, 0}
	modifySlice(s2)
	fmt.Printf("%#v", s2) //[]int{1, 0, 0}

	// map
	m2 := make(map[int]string)
	fmt.Printf("m2 is not nill --> %#v \n ", m2) //map[int]string{}
	modifyMap(m2)
	fmt.Printf("m2 is not nill --> %#v \n ", m2) // map[int]string{0:"string"}

	// chan
	c2 := make(chan string)
	fmt.Printf("c2 is not nill --> %#v \n ", c2) // (chan string)(0xc42007c060)
	go modifyChan(c2)
	fmt.Printf("c2 is not nill --> %#v ", <-c2) //"string"
}

func modifySlice(s []int) {
	s[0] = 1
}

func modifyMap(m map[int]string) {
	m[0] = "string"
}
func modifyChan(c chan string) {
	c <- "string"
}
