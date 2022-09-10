package main

import "fmt"

func main()  {
	// [...]type 是声明不定长的 array, 但终究数组长度是要在运行中确定下来的, 并不意味着这是在声明一个 slice
	// []type 才是在声明 slice

	// array
	a := [...]string{"zhang", "lin", "wang"}
	b := [...]string{"zhang", "lin", "wang"}
	c := [...]string{"zhang", "lin", "guan"}

	fmt.Println(equal(a[:], b[:]))		// true
	fmt.Println(equal(a[:], c[:]))		// false
	fmt.Println(equal(a[:2], c[:2]))	// true

	// slice
	x := []string{"zhang", "lin", "wang"}
	y := []string{"zhang", "lin", "wang"}
	z := []string{"zhang", "lin", "guan"}

	fmt.Println(equal(x, y))		// true
	fmt.Println(equal(x, z))		// false
}

func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}

	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}

	return true
}
