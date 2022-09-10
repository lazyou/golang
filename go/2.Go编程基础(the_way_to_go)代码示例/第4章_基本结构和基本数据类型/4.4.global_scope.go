package main

var a = "G"

func main() {
	n()
	m()
	n()
}

func n() {
	print(a)
}

func m() {
	// 全局变量 a 的值被覆盖
	a = "O"
	print(a)
}
