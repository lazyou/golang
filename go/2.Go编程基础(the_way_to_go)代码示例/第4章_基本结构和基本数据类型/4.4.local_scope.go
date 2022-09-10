package main

var a = "G"

func main() {
	n() // G
	m() // O
	n() // G
}

func n() {
	print(a)
}

func m() {
	// 此处 a 是 m() 函数内的作用域，外层的 var a 没有起作用
	a := "O"
	print(a)
}
