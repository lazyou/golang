package main

import (
	"fmt"
)

func main() {
	var s1 []int

	if s1 == nil {
		// s1 is nil --> []int(nil)
		fmt.Printf("s1 is nil --> %#v \n", s1)
	}

	// slice
	s2 := make([]int, 3)

	if s2 == nil {
		fmt.Printf("s2 is nil --> %#v \n ", s2)
	} else {
		// s2 is not nill --> []int{0, 0, 0}
		fmt.Printf("s2 is not nill --> %#v \n ", s2)
	}
}
