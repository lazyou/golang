package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("In main()")

	// 协程
	go longWait()
	go shortWait()

	fmt.Println("About to sleep in main()")

	// sleep works with a Duration in nanoseconds (ns) !
	// 睡眠 10秒
	time.Sleep(10 * 1e9)
	fmt.Println("At the end of main()")
}

func longWait() {
	fmt.Println("longWait() Beginning")
	time.Sleep(5 * 1e9) // sleep for 5 seconds
	fmt.Println("longWait() End")
}

func shortWait() {
	fmt.Println("shortWait() Beginning")
	time.Sleep(2 * 1e9) // sleep for 2 seconds
	fmt.Println("shortWait() End")
}
