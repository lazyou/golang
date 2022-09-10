package main

import (
	"flag" // command line option parser
	"os"
)

var NewLine = flag.Bool("n", false, "print newline") // echo -n flag, of type *bool

const (
	Space   = " "
	Newline = "\n"
)

// 命令行 选项参数 获取
// go run command_flag.go A B C 输出 "A B C"
// go run command_flag.go -n A B C 输出 "A 换行 B 换行 C 换行"
func main() {
	// 默认打印选项提示
	flag.PrintDefaults()

	// 扫描参数列表（或者常量列表）并设置 flag
	flag.Parse() // Scans the arg list and sets up flags

	var s string = ""
	// flag.Narg() 返回参数的数量
	for i := 0; i < flag.NArg(); i++ {
		if i > 0 {
			s += Space
			if *NewLine { // -n is parsed, flag becomes true
				s += Newline
			}
		}

		// flag.Arg(i) 表示第i个参数
		// flag.Arg(0) 就是第一个真实的 flag，而不是像 os.Args(0) 放置程序的名字
		s += flag.Arg(i)
	}

	os.Stdout.WriteString(s)
}