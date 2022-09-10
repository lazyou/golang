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
	// inputReader := bufio.NewReader(os.Stdin) 这行代码，将会创建一个读取器，并将其与标准输入绑定
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input: ")
	input, err = inputReader.ReadString('\n')
	if err == nil {
		fmt.Printf("The input was: %s\n", input)
	}
}
