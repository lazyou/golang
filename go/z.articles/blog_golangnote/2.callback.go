package main

import (
	"fmt"
	"strconv"
)

// 一个记录日志的类型：func(string)
type saveLog func(msg string)

// 将字符串转换为 int64, 如果转换失败调用 saveLog
func stringToInt(s string, log saveLog) int64 {
	if value, err := strconv.ParseInt(s, 0, 0); err != nil {
		log(err.Error())

		return 0
	} else {
		return value
	}
}

// 记录日志消息的具体实现
func myLog(msg string) {
	fmt.Println("Find Error:", msg)
}

func main() {
	fmt.Println(stringToInt("123", myLog))

	fmt.Println(stringToInt("a", myLog))
}
