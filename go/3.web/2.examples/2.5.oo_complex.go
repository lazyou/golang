package main

import (
	"fmt"
)

/*
https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/02.5.md#指针作为receiver

TODO: 指针作为receiver

现在让我们回过头来看看 SetColor 这个 method，它的 receiver 是一个指向 Box 的指针，是的，你可以使用 *Box 。想想为啥要使用指针而不是Box本身呢？

我们定义SetColor的真正目的是想改变这个Box的颜色，如果不传Box的指针，那么SetColor接受的其实是Box的一个copy，也就是说method内对于颜色值的修改，其实只作用于Box的copy，而不是真正的Box。所以我们需要传入指针。

这里可以把receiver当作method的第一个参数来看，然后结合前面函数讲解的传值和传引用就不难理解

这里你也许会问了那 SetColor 函数里面应该这样定义 *b.Color=c,而不是 b.Color=c,因为我们需要读取到指针相应的值。

你是对的，其实Go里面这两种方式都是正确的，当你用指针去访问相应的字段时(虽然指针没有任何的字段)，Go知道你要通过指针去获取这个值，看到了吧，Go的设计是不是越来越吸引你了。

也许细心的读者会问这样的问题，PaintItBlack里面调用SetColor的时候是不是应该写成(&bl[i]).SetColor(BLACK)，因为SetColor的receiver是*Box，而不是Box。

你又说对的，这两种方式都可以，因为Go知道receiver是指针，他自动帮你转了。

__也就是说：__
	如果一个method的receiver是*T,你可以在一个T类型的实例变量V上面调用这个method，而不需要&V去调用这个method
类似的

	如果一个method的receiver是T，你可以在一个*T类型的变量P上面调用这个method，而不需要 *P去调用这个method
所以，你不用担心你是调用的指针的method还是不是指针的method，Go知道你要做的一切，这对于有多年C/C++编程经验的同学来说，真是解决了一个很大的痛苦。
*/

const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

type Color byte

type Box struct {
	width, height, depth float64
	color                Color
}

type BoxList []Box // a slice of boxes

// 体积
func (b Box) Volume() float64 {
	return b.width * b.height * b.depth
}

// 指针类型
func (b *Box) SetColor(c Color) {
	b.color = c
	// *b.color = c // TODO: 为什么这样写不行?
}

// 找出最大 Box 的颜色
func (bl BoxList) BiggestColor() Color {
	v := 0.00
	k := Color(WHITE)

	// 遍历 Box 数组
	for _, b := range bl {
		if bv := b.Volume(); bv > v {
			v = bv
			k = b.color
		}
	}

	return k
}

func (bl BoxList) PaintItBlack() {
	for i, _ := range bl {
		bl[i].SetColor(BLACK)
	}
}

func (c Color) String() string {
	strings := []string{"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}

	return strings[c]
}

func main() {
	// TODO: 这种结构似乎很像数据库查出来的结果
	boxes := BoxList{
		Box{4, 4, 4, RED},
		Box{10, 10, 1, YELLOW},
		Box{1, 1, 20, BLACK},
		Box{10, 10, 1, BLUE},
		Box{10, 30, 1, WHITE},
		Box{20, 20, 20, YELLOW},
	}

	// 盒子数量
	fmt.Printf("We have %d boxes in our set\n", len(boxes))

	fmt.Println(boxes)

	// 第一个 盒子 的体积
	fmt.Println("The volume of the first one is", boxes[0].Volume(), "cm³")

	// 最后一个 盒子 的颜色 (链式调用)
	fmt.Println("The color of the last one is", boxes[len(boxes)-1].color.String())

	// 最大的盒子 的 颜色
	fmt.Println("The biggest one is", boxes.BiggestColor().String())

	fmt.Println("Let's paint them all black")

	// 全部盒子设置位 黑色
	boxes.PaintItBlack()

	fmt.Println("The color of the second one is", boxes[1].color.String())

	fmt.Println("Obviously, now, the biggest one is", boxes.BiggestColor().String())
}
