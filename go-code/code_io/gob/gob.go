package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

type P struct {
	X, Y, Z int
	Name    string
}

type Q struct {
	X, Y *int32
	Name string
}

// Gob 是 Go 自己的以二进制形式序列化和反序列化程序数据的格式
func main()  {
	var network bytes.Buffer
	enc := gob.NewEncoder(&network) // Will write to network.
	dec := gob.NewDecoder(&network) // Will read from network.

	// gob 编码: 只有可导出的字段会被编码，零值会被忽略
	err := enc.Encode(P{3, 4, 5, "Pythagoras"})
	if err != nil {
		log.Fatal("encode error:", err)
	}

	// gob 解码: 数据结构不对应
	var q Q
	err = dec.Decode(&q)
	if err != nil {
		log.Fatal("decode error:", err)
	}
	fmt.Printf("%q: {%d,%d}\n", q.Name, q.X, q.Y)
}
