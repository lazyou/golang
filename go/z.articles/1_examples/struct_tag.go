package main

import (
	"encoding/json"
	"fmt"
)

// struct tag: 这种方式主要是用在xml，json和struct间相互转换，非常方便直观，
// 比如接口给的参数一般是json传过来，但是内部我们要转为struct再进行处理
type User struct {
	Name string `json:"userName"`
	Age  int    `json:"userAge"`
}

func main() {
	var user User
	user.Name = "nick"
	user.Age = 18
	conJson, waht := json.Marshal(user)

	// {"userName":"nick","userAge":18} TODO: 注意这里 key 值变了, 这是 tag 的作用之一
	fmt.Println(string(conJson), waht)
}
