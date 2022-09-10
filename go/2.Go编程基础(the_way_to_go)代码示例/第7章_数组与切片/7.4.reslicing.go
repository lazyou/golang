package main

import "fmt"

// 改变切片长度的过程称之为切片重组 reslicing，做法如下：slice1 = slice1[0:end]
func main() {
	slice1 := make([]int, 0, 10)

	// load the slice, cap(slice1) is 10:
	for i := 0; i < cap(slice1); i++ {
		slice1 = slice1[0 : i+1]
		slice1[i] = i
		fmt.Printf("The length of slice is %d\n", len(slice1))
	}

	// print the slice:
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
}
