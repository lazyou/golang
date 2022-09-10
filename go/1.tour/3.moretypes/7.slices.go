package main

import "fmt"

/*
slice: 这就是灵活的 array?
	一个 slice 会指向一个序列的值，并且包含了长度信息。

	[]T 是一个元素类型为 T 的 slice。

	len(s) 返回 slice s 的长度。
*/
func main()  {
	s := []int{2, 3, 5, 7, 11, 13}
	// s == [2 3 5 7 11 13]
	fmt.Println("s ==", s)

	/*
	s[0] == 2s[1] == 3
	s[2] == 5s[3] == 7
	s[4] == 11
	s[5] == 13	
	*/
	for i := 0; i < len(s); i++ {
		fmt.Printf("s[%d] == %d\n", i, s[i])
	}
}