package main

import "fmt"

type User struct {
	Name string
	Age  int32
	mess string
}

// 这是一种“继承”的写法，在go语言中这种方式叫做“嵌入”（embed）
type Player struct {
	User
}

func main() {
	user := User{"lin", 99, "mess"}

	fmt.Printf("Type: %T Value: %v\n", user, user)

	// Player 结构体继承来自 User 结构体的属性
	player := Player{}
	player.Name = "ln"
	player.Age = 22
	fmt.Printf("Type: %T Value: %v\n", player, player)
}
