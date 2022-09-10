package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	// Printf: %T 格式化输出该值的类型; %v 输出为相应值的默认格式
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)

	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)

	fmt.Printf("Type: %T Value: %v\n", z, z)
}
