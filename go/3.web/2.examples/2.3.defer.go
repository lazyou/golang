package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func deferExample1() {
	inputFile := "defer.txt"
	// inputFile := "defer_not_exist.txt"

	fi, err := os.Open(inputFile)

	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
		// panic(err.Error())
	}

	// defer是采用后进先出模式
	defer fi.Close()

	// 逐行读取文件内容
	br := bufio.NewReader(fi)

	for {
		a, _, c := br.ReadLine()

		if c == io.EOF {
			break
		}

		fmt.Println(string(a))
	}
}

func deferExample2() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

func main() {
	deferExample1()

	// defer是采用后进先出模式，所以如下代码会输出4 3 2 1 0
	deferExample2()
}
