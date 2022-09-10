package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var col1, col2, col3 []string
	for {
		var v1, v2, v3 string

		// 空格分隔, 分别读取到变量中
		_, err := fmt.Fscanln(file, &v1, &v2, &v3)
		if err == io.EOF {
			break
		}

		//fmt.Println(v1, v2, v3)

		col1 = append(col1, v1)
		col2 = append(col2, v2)
		col3 = append(col3, v3)
	}

	fmt.Println(col1)
	fmt.Println(col2)
	fmt.Println(col3)
}
