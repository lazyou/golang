package main

import "fmt"

// 2.2 切片 slice : 和声明array一样，只是少了长度
// slice 并不是真正意义上的动态数组，而是一个 __引用类型__
func main()  {
	// 和声明 array 一样, 只是少了长度
	//var fslice []int

	slice := []byte {'a', 'b', 'c'}
	fmt.Println(slice) // [97 98 99]
	fmt.Println(len(slice), cap(slice))

	// slice 可以从一个数组或一个已经存在的slice中再次声明
	// slice 通过 array[i:j] 来获取，其中 i 是数组的开始位置，j 是结束位置，但不包含 array[j]，它的长度是 j-i
	var ar = [10]byte {'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	var a, b []byte

	a = ar[2:5]
	b = ar[3:5]

	fmt.Println(a)
	fmt.Println(b)

	/**
	slice 便捷操作:
		slice 的默认开始位置是0， ar[:n] 等价于 ar[0:n];

		slice 的第二个序列默认是数组的长度， ar[n:] 等价于 ar[n:len(ar)];

		如果从一个数组里面直接获取slice，可以这样 ar[:] ，因为默认第一个序列是 0，第二个是数组的长度，即等价于 ar[0:len(ar)].
	 */

	 /**
	slice 几个内置函数:
		len 获取 slice 的长度;

		cap 获取 slice 的最大容量;

		append 向 slice 里面追加一个或者多个元素，然后返回一个和 slice 一样类型的 slice;

		copy 函数 copy 从源 slice 的 src 中复制元素到目标 dst，并且返回复制的元素的个数
	  */
}
