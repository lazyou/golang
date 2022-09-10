package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	time.Sleep(100 * time.Millisecond)
	end := time.Now()

	fmt.Printf("longCalculation took this amount of time: %s\n", end.Sub(start))
}
