package main

import (
	"fmt"
)

// var 语句用于声明一个变量列表: 类型在最后
var c, python, java bool

func main()  {
    // 鼠标放上去是提示警告信息是“should omit type int from declaration of var v1; it will be inferred from the right-hand side”。
    // 意思是应该省略VAR V1的类型为int;它将从右侧来推断。可以看出指定类型已不再是必需的，Go可以从初始化表达式的右值推导出该变量类型
    var i int = 1
    
    // 没有赋值默认是 零值
    fmt.Println(i, c, python, java)
}
