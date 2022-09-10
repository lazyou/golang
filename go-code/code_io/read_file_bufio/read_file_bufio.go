package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main()  {
	inputFile, inputErr := os.Open("./input.dat")
	if inputErr != nil {
		fmt.Println("文件读取出错")
		fmt.Println(inputErr.Error())
		return
	}
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readerError := inputReader.ReadString('\n')
		fmt.Printf("The input was: %s", inputString)
		if readerError == io.EOF {
			return
		}
	}
}
