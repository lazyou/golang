package main

import (
	"fmt"
	"strconv"
)

func main() {
	print(StrToInt("999"))

	print(StrToInt64("999"))

	print(IntToStr(1122))

	print(Int64ToStr(1122))
}

func print(v interface{})  {
	fmt.Printf("%#v \n", v)
}

// StrToInt string 转 int
func StrToInt(str string) int {
	result, _ := strconv.Atoi(str)

	return result
}

// StrToInt64 string 转 int64
func StrToInt64(str string) int64 {
	result, _ := strconv.ParseInt(str, 10, 64)

	return result
}

// IntToStr int 转 string
func IntToStr(i int) string {
	result := strconv.Itoa(i)

	return result
}

// Int64ToStr int64 转 string
func Int64ToStr(i int64) string {
	result := strconv.FormatInt(i, 10)

	return result
}