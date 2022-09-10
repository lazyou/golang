package main

import (
	"os"
)

func test() error {
	f, err := os.Create("test.txt")

	if err != nil {
		return err
	}

	// 注册调用，而不是注册函数。必须提供参数，哪怕为空
	defer f.Close()

	f.WriteString("Hello, World!")

	return nil
}

func main() {
	test()
}
