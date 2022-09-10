package main

import "fmt"

type struct1 struct {
	i1  int
	f1  float32
	str string
}

func main() {
	// 声明 1
	ms := new(struct1)
	fmt.Println(ms) // &{0 0 } new 得到的是指针

	ms.i1 = 10
	ms.f1 = 15.5
	ms.str = "Chris"
	fmt.Println(ms)

	// 声明 2
	var ms2 struct1
	ms2 = struct1{10, 15.5, "Chris"}
	fmt.Println(ms2) // {10 15.5 Chris}
}
