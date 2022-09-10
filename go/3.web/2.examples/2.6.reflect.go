package main

import (
	"fmt"
	"reflect"
)

func main() {
	i := 1

	// 得到类型的元数据,通过 t 我们能获取类型定义里面的所有元素
	t := reflect.TypeOf(i)

	//得到实际的值，通过v我们获取存储在里面的值，还可以去改变值
	v := reflect.ValueOf(i)

	fmt.Println(t, v) // int, 1

	// 通过反射解析 struct
	type Todo struct {
		Id    int64  `orm:"auto"`
		Title string `orm:"size(128)"`
		Sort  int
	}

	todo := Todo{
		Id:    1,
		Title: "标题title",
		Sort:  0,
	}

	todoType := reflect.TypeOf(todo) // 转化为reflect对象
	todoValue := reflect.ValueOf(todo)
	fmt.Println(todoType, todoValue)
	fmt.Println()

	fmt.Println(todoType.Name())  // `Todo`
	fmt.Println(todoType.Size())  // 32
	fmt.Println(todoValue.Type()) // main.Todo
	fmt.Println(todoValue.Kind()) // struct
	fmt.Println()

	fmt.Println(todoType.Field(0).Name)   // Id
	fmt.Println(todoType.Field(0).Type)   // int64
	fmt.Println(todoType.Field(0).Tag)    // orm:"auto" -- 标签
	fmt.Println(todoValue.Field(0).Int()) // 1 -- 这里只有选择对类型才行
	fmt.Println()

	fmt.Println(todoType.Field(1).Name)      // Title
	fmt.Println(todoType.Field(1).Type)      // string
	fmt.Println(todoType.Field(1).Tag)       // orm:"size(128)"
	fmt.Println(todoValue.Field(1).String()) // 标题title -- 这里只有选择对类型才行
	fmt.Println()

	// 通过反射修改值 -- TODO: 会报错哇， 不知道怎么修改值
	//todoValue.Field(0).SetInt(998)
	//todoValue.Field(1).SetString("通过反射改变value")
	//fmt.Println(todo)

	// 反射修改值
	x := 3.4
	fmt.Println("修改前:", x)
	reflect.ValueOf(&x).Elem().SetFloat(7.1) // __必须__ 传入指针, 且调用 Elem() 方法
	fmt.Println("修改后:", x)
}
