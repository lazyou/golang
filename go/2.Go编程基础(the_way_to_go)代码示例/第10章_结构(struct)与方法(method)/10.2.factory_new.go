// 首字母小写，私有化
type matrix struct {
	id   int64
	age  int64
	name string
}

func NewMatrix(params) *matrix {
	m := new(matrix) // 初始化 m
	return m
}


// 在其他包里使用
package main
import "matrix"
...
wrong := new(matrix.matrix)     // 编译失败（matrix 是私有的）
right := matrix.NewMatrix(...)  // 实例化 matrix 的唯一方式
