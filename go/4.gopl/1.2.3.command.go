package main

import (
	"fmt"
	"strings"
	"os"
)

// `go run 1.2.3.command.go canshu1 canshu2`
func main() {
	fmt.Println(strings.Join(os.Args[0:], " "))
}