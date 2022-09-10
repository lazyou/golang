// Interface嵌套
package main

type Car interface {
	NameGet() string
	Run(n int)
	Stop()
}

// 所以想要实现 Used 接口, 也必须实现 Car 接口
type Used interface {
	// 嵌套了 Car 接口
	Car
	Cheap()
}
