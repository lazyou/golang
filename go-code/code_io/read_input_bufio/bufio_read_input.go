package main

import (
	"bufio"
	"fmt"
	"os"
)

var inputReader *bufio.Reader
var input string
var err error

func main() {
	fmt.Println("Please enter some input: ")
	// 用 bufio 包提供的缓冲读取（buffered reader）来读取数据
	inputReader = bufio.NewReader(os.Stdin)
	input, err = inputReader.ReadString('\n')
	if err == nil {
		fmt.Printf("The input was: %s \n", input)
	}
}

