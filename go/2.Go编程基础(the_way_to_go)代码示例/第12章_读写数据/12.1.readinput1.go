// 从控制台读取输入:
package main

import "fmt"

var (
	firstName, lastName, s string
	i                      int
	f                      float32
	input                  = "56.12 / 5212 / Go"
	format                 = "%f / %d / %s"
)

func main() {
	// Scanln 扫描来自标准输入的文本，将 [空格分隔] 的值依次存放到后续的参数内，直到碰到换行
	fmt.Println("Please enter your full name: ")
	fmt.Scanln(&firstName, &lastName)
	fmt.Printf("Hi %s %s!\n", firstName, lastName)

	// Scanf 的第一个参数用作格式字符串，用来决定如何读取
	fmt.Sscanf(input, format, &f, &i, &s)
	fmt.Println("From the string we read: ", f, i, s)
	// 输出结果: From the string we read: 56.12 5212 Go
}
