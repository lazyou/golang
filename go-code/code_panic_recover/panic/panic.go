package main

import "fmt"

// 当错误条件（我们所测试的代码）很严苛且不可恢复，程序不能继续运行时，可以使用 panic 函数产生一个中止程序的运行时错误
func main() {
	fmt.Println("Starting the program")
	panic("A severe error occurred: stopping the program!")
	fmt.Println("Ending the program")
}
