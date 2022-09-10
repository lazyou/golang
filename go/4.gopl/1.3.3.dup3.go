package main

import (
	"os"
	"io/ioutil"
	"fmt"
	"strings"
)

// 实现上，bufio.Scanner、ioutil.ReadFile 和 ioutil.WriteFile 都使用 *os.File 的 Read 和 Write 方法，
// 但是，大多数程序员很少需要直接调用那些低级（lower-level）函数。高级（higher-level）函数，像 bufio和io/ioutil 包中所提供的那些，用起来要容易点。
func main() {
	counts := make(map[string]int)

	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)

		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}

		// ReadFile函数返回一个字节切片（byte slice）
		// 必须把它转换为string，才能用strings.Split分割
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	for line, n := range counts {
		fmt.Printf("%d\t%s\n", n, line)
	}
}