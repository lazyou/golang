package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	outputFile, outputError := os.OpenFile("./output.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if outputError != nil {
		fmt.Println(outputError.Error())
		return
	}
	defer outputFile.Close()

	outputWriter := bufio.NewWriter(outputFile)
	outputString := "hello world!\n"

	for i := 0; i < 10; i++ {
		outputWriter.WriteString(outputString)
	}

	// 缓冲区的内容紧接着被完全写入文件
	_ = outputWriter.Flush()
}