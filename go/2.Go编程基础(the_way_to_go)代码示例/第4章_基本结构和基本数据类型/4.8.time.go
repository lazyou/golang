package main

import (
	"fmt"
	"time"
)

var week time.Duration

func main() {
	t := time.Now()
	fmt.Println(t)
	fmt.Printf("%4d-%02d-%02d\n", t.Year(), t.Month(), t.Day())

	t = time.Now().UTC()
	fmt.Println(t)
}
