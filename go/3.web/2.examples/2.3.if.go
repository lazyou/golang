package main

import "fmt"

func main() {
	// 1.
	x := 1

	if x > 10 {
		fmt.Println("x is greater than 10")
	} else {
		fmt.Println("x is less than 10")
	}

	// 2. 条件判断语句里面允许声明一个变量，这个变量的作用域只能在该条件逻辑块内，其他地方就不起作用了
	if y := getY(); y > 10 {
		fmt.Println("y is greater than 10")
	} else {
		fmt.Println("y is less than 10")
	}

	// if ... else if ... else ...
	integer := 4
	if integer == 3 {
		fmt.Println("The integer is equal to 3")
	} else if integer < 3 {
		fmt.Println("The integer is less than 3")
	} else {
		fmt.Println("The integer is greater than 3")
	}
}

func getY() int {
	return 11
}
