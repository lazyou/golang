package main

import "fmt"

var (
	firstName, lastName, s string
	i int
	f float32
	input = "56.12 / 5212 / Go"
	format = "%f / %d / %s"
)

func main() {
	// 从输入中读取
	fmt.Println("Please enter your full name:")
	fmt.Scanln(&firstName, &lastName)
	fmt.Printf("Hi %s %s! \n", firstName, lastName)

	// 从字符串中读取
	fmt.Sscanf(input, format, &f, &i, &s)
	fmt.Println("From the string we read:", f, i, s)
}
