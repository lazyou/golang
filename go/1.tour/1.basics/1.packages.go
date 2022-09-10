// 程序运行的入口是包 main
package main

// 这个程序使用并导入了包 "fmt" 和 "math/rand"
import (
	"fmt"
	"math/rand"
)

func main() {
	// 这个程序的运行环境是确定性的，因此 rand.Intn 每次都会返回相同的数字
	fmt.Println("My favorite number is", rand.Intn(10))
}
