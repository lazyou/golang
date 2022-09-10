package main

import (
	"fmt"
)

// 函数支持多个返回值
func SumAndProduct(A, B int) (int, int) {
	return A + B, A * B
}

// 另一种多个返回值的写法
func SumAndProductAnothor(A, B int) (sum int, product int) {
	sum = A + B
	product = A * B
	return  // 此处省略返回值, 因为在函数定义事已经声明了
}

func main() {
	x := 3
	y := 4

	result1, result2 := SumAndProduct(x, y)
	fmt.Println(result1)
	fmt.Println(result2)

	sum, product := SumAndProductAnothor(x, y)
	fmt.Println(sum)
	fmt.Println(product)
}