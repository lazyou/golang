package main

import "fmt"

func main() {
	var m1 map[int]string
	// map 零值应该是 nil
	if m1 == nil {
		fmt.Printf("m1 is nil --> %#v \n ", m1) //map[int]string(nil)
	}

	m2 := make(map[int]string)
	if m2 == nil {
		fmt.Printf("m2 is nil --> %#v \n ", m2)
	} else {
		// make 产生的 map 不再是零值
		fmt.Printf("m2 is not nill --> %#v \n ", m2) // map[int]string{}
	}
}
