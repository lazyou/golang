package main

import (
	"fmt"
)

func main() {
	bool1 := true
	bool2 := true

	// if 里面的条件不需要使用括号包裹
	if bool1 && bool2 {
		fmt.Printf("The value is true\n")
	} else {
		fmt.Printf("The value is false\n")
	}

	// If 的简短方式
	if bool3 := true; bool3 {
		fmt.Println("If 的简短方式")
	}
}
