// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 179.

// The 7.4.sleep program sleeps for a specified period of time.
package main

import (
	"flag"
	"fmt"
	"time"
)

//!+7.4.sleep
var period = flag.Duration("period", 1*time.Second, "7.4.sleep period")

// go run 7.4.sleep/sleep.go -period 10000ms
func main() {
	flag.Parse()
	fmt.Printf("Sleeping for %v...", *period)
	time.Sleep(*period)
	fmt.Println()
}

//!-7.4.sleep
