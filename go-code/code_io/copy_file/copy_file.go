package main

import (
	"fmt"
	"io"
	"os"
)

func main()  {
	i, err := CopyFile("target.txt", "source.txt")
	if err != nil {
		fmt.Println("Copy error!", i)
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Copy done!", i)
}

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}
