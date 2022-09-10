package main

import "fmt"

// 2.2 字符串 string: 声明形式("" 和 ``), 不可变, 及其修改
func main()  {
	// 多行字符串声明 (Raw) ``
	m := `hello
			world`
	fmt.Println(m)

	// 如何修改字符串
	s :="hello" // 如果是中文就会出现乱码
	c :=[]byte(s) // 将字符串 s 转换为 []byte 类型
	c[0] = 'c'
	s2 := string(c) // 再转换回 string 类型
	fmt.Println(s2)

	// 连接字符串 +
	h :="hello"
	w := " world"
	hw := h + w
	fmt.Println(hw)
}
