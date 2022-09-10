package main

import (
	"os"
	"fmt"
)

// `go run 1.2.2.command.go canshu1 canshu2`
func main() {
	s, sep := "", ""

	for _, arg := range os.Args[0:] {
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)
}
