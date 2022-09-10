package main

import (
	"fmt"

	"./trans"
)

// cannot find package "./trans" in:
// TODO: 如何引用本地相对路径的包？
var twoPi = 2 * trans.Pi

func main() {
	fmt.Printf("2*Pi = %g\n", twoPi)
}
