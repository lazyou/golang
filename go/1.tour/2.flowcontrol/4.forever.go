package main

/*
ctrl + c 推出死循环
```
/mnt/e/codes/go/tour/2.flowcontrol$ go run 4.forever.go
^Csignal: interrupt  // 居然会输出这个
```
*/
func main()  {
	// 死循环: 如果省略了循环条件，循环就不会结束，因此可以用更简洁地形式表达死循环
	for {
	}
}