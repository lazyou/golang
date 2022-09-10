package main

import "fmt"

func main() {
	// 切片 结构体
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
	}

	// Type: []struct { i int; b bool } Value: [{2 true} {3 false}]
	fmt.Printf("Type: %T Value: %v\n", s, s)
}
