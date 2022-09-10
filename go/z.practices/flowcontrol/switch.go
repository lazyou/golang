package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")

	// TODO: 缩进方式有点难受. 不习惯
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.", os)
	}
}
