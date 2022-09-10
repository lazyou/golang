package main

import (
	"fmt"
)

type User struct{
    Name string
    Age int32
    mess string
}


func main()  {
    var user1 User          // 普通结构体, 非指针类型
    var user2 = &User{}     // 实际上转化为 var user1 *User = &User{}       -- 所以这里是指针类型
    var user3 = new(User)   // 实际上转化为 var user2 *User = new(User)     -- 所以这里是也是指针类型

    user1.Name = "user1"
    user2.Name = "user2"
    user3.Name = "user3"

    // main.User {user1 0 }     -- 非指针
    fmt.Printf("%T %v\n", user1, user1)

    // *main.User &{user2 0 }   -- 指针
    fmt.Printf("%T %v\n", user2, user2)

    // *main.User &{user2 0 }   -- 指针
    fmt.Printf("%T %v\n", user3, user3)
}
