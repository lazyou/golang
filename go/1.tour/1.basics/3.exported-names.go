package main

import (
	"fmt"
	"math"
)

func main() {
	// 在 Go 中，首字母大写的名称是被导出的
	fmt.Println(math.Pi)

	// fmt.Println(math.pi) 这句是错的, pi 首字母没大写
}
