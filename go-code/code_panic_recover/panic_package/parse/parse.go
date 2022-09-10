package parse

import (
	"fmt"
	"strconv"
	"strings"
)

// 自定义包中的错误处理和 panicking

//1）在包内部，总是应该从 panic 中 recover：不允许显式的超出包范围的 panic()
//2）向包的调用者返回错误值（而不是 panic）。

// 在包内部，特别是在非导出函数中有很深层次的嵌套调用时，将 panic 转换成 error 来告诉调用方为何出错，是很实用的（且提高了代码可读性）。

// ParseError表示将单词转换为整数时出错。
type ParseError struct {
	Index int 		//空格分隔的单词列表的索引。
	Word string 	// 产生解析错误的单词。
	Err error 		//导致此错误的原始错误（如果有）。
}

// 返回可读的错误消息
func (e *ParseError) String() string {
	return fmt.Sprintf("pkg parse: error parsing %q as int", e.Word)
}

func fields2numbers(fields []string) (numbers []int) {
	if len(fields) == 0 {
		panic("no words to parse")
	}

	for idx, field := range fields {
		num, err := strconv.Atoi(field)
		if err != nil {
			panic(&ParseError{idx, field, err})
		}

		numbers = append(numbers, num)
	}

	return
}

// Parse 将输入中以空格分隔的单词解析为整数。
func Parse(input string) (numbers []int, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("pkg: %v", r)
			}
		}
	}()

	fields := strings.Fields(input)
	numbers = fields2numbers(fields)
	return
}
