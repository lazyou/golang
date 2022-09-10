package main

import (
	"fmt"
	"math/cmplx"
)

/*
Go 的基本类型有Basic types:
```go
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // uint8 的别名

rune // int32 的别名
     // 代表一个Unicode码

float32 float64

complex64 complex128
```

// int，uint 和 uintptr 类型在32位的系统上一般是32位，而在64位系统上是64位
*/

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	const f = "%T(%v)\n"
	// bool(false)
	fmt.Printf(f, ToBe, ToBe)

	// uint64(18446744073709551615)
	fmt.Printf(f, MaxInt, MaxInt)

	// complex128((2+3i))
	fmt.Printf(f, z, z)
}
