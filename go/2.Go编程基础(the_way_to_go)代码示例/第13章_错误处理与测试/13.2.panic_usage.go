// 当发生错误必须中止程序时，panic 可以用于错误处理模式：
if err != nil {
	panic("ERROR occurred:" + err.Error())
}