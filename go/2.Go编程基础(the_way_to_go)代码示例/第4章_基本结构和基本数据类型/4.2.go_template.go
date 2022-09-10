package main

// Go 程序的一般结构

import (
	"fmt"
)

const c = "C"

var v int = 5

type T struct{}

// initialization of package
func init() {
	fmt.Println("init")
}

func main() {
	var a int
	Func1()
	fmt.Println(a)
}

func (t T) Method1() {
	fmt.Println("Method1")
}

func Func1() {
	fmt.Println("Func1")
}
