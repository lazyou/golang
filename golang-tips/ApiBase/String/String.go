package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// TODO: 如何在里面放变量?
	str := "Hello World! \nHello Gopher! \n"
	fmt.Println(str)

	// 内置的len()函数获取的是每个字符的UTF-8编码的长度和
	s := "其实就是rune"
	fmt.Println(len(s))                    // 16
	fmt.Println(utf8.RuneCountInString(s)) // 8 (通常这个才是我们想要的字符串长度)

	// 如字符串含有中文等字符，我们可以看到每个中文字符的索引值相差 3
	ss := "Go语言四十二章经"
	for k, v := range ss {
		fmt.Printf("index：%d, value：%c == %d\n", k, v, v)
	}

	// 更多字符串操作
	/*
		标准库中有四个包对字符串处理尤为重要：bytes、strings、strconv和unicode包。

		strings包提供了许多如字符串的查询、替换、比较、截断、拆分和合并等功能。

		bytes包也提供了很多类似功能的函数，但是针对和字符串有着相同结构的[]byte类型。因为字符串是只读的，因此逐步构建字符串会导致很多分配和复制。在这种情况下，使用bytes.Buffer类型将会更有效，稍后我们将展示。

		strconv包提供了布尔型、整型数、浮点数和对应字符串的相互转换，还提供了双引号转义相关的转换。

		unicode包提供了IsDigit、IsLetter、IsUpper和IsLower等类似功能，它们用于给字符分类。

		strings 包提供了很多操作字符串的简单函数，通常一般的字符串操作需求都可以在这个包中找到。下面简单举几个例子：

		判断是否以某字符串打头/结尾 strings.HasPrefix(s, prefix string) bool strings.HasSuffix(s, suffix string) bool

		字符串分割 strings.Split(s, sep string) []string

		返回子串索引 strings.Index(s, substr string) int strings.LastIndex 最后一个匹配索引

		字符串连接 strings.Join(a []string, sep string) string 另外可以直接使用“+”来连接两个字符串

		字符串替换 strings.Replace(s, old, new string, n int) string

		字符串转化为大小写 strings.ToUpper(s string) string strings.ToLower(s string) string

		统计某个字符在字符串出现的次数 strings.Count(s, substr string) int

		判断字符串的包含关系 strings.Contains(s, substr string) bool
	*/
}
