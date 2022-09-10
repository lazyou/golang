package main

import (
	"fmt"
	"image"
)

/*
图片
	Package image 定义了 Image 接口：

	package image

	type Image interface {
		ColorModel() color.Model
		Bounds() Rectangle
		At(x, y int) color.Color
	}
	
	注意：Bounds 方法的 Rectangle 返回值实际上是一个 image.Rectangle， 其定义在 image 包中
*/
func main() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}