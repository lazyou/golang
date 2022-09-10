package main

import (
	"archive/zip"
	"log"
	"os"
)

/**
https://golang.org/pkg/archive/zip

Package zip provides support for reading and writing ZIP archives.

为了向后兼容，文件头包含32位和64位字段。64位字段总是包含正确的值，对于普通存档，这两个字段是相同的。对于需要ZIP64格式的文件，32位字段将是0xffffffff，必须使用64位字段。
*/

// https://golang.org/src/archive/zip/example_test.go
func main() {
	ExampleWriter()
}

// 官方案例没看懂, 稍作修改: 创建一个 zip 包并且写入文件
func ExampleWriter() {
	// Create a new zip archive.
	z, _ := os.Create("ExampleWriter.zip")
	w := zip.NewWriter(z)
	defer w.Close()

	// Add some files to the archive.
	var files = []struct {
		Name, Body string
	}{
		{"readme.txt", "This archive contains some text files."},
		{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "Get animal handling licence.\nWrite more examples."},
	}

	for _, file := range files {
		f, err := w.Create(file.Name)
		checkError(err)

		_, err = f.Write([]byte(file.Body)) // string 转为 byte 才能被写入
		checkError(err)
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
