package main

import "fmt"

var arr1 [5]int

func main() {
	// for 赋值数组 打印数组
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i * 2
	}

	for i := 0; i < len(arr1); i++ {
		fmt.Printf("Array at index %d is %d\n", i, arr1[i])
	}

	// for range
	for key, val := range arr1 {
		fmt.Printf("Array at index %d is %d\n", key, val)
	}

	// 切片
	// a := []string{"a", "b", "c", "d"}
	a := [...]string{"a", "b", "c", "d"} // 三个点可以省略
	for i := range a {
		fmt.Println("Array item", i, "is", a[i])
	}

	// 数组引用传递
	var ar [3]int
	f(ar)   // passes a copy of ar
	fp(&ar) // passes a pointer to ar
}

func f(a [3]int)   { fmt.Println(a) }
func fp(a *[3]int) { fmt.Println(a) } // &[0 0 0]
