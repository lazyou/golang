package main

import "fmt"

func main() {
	mf := map[int]func() int{
		1: func() int { return 10 },
		2: func() int { return 20 },
		5: func() int { return 50 },
	}

	fmt.Println(mf)

	f := func() int { return 10 }
	fmt.Println(f) // 0x484e60
}
