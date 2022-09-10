package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

var ch chan int
var nickname string

// 当前连接
var currentConn *net.TCPConn

// 读取行文本
var inputText = bufio.NewReader(os.Stdin)

// reader 用户输入内容读取
// 这里的 conn 是同一个 TCPConn
// 每个 Client 都不断地读取里面内容
func reader(conn *net.TCPConn) {
	buff := make([]byte, 128)

	for {
		j, err := conn.Read(buff)
		if err != nil {
			ch <- 1
			break
		}

		fmt.Println(string(buff[0:j]), " -- from reader")
	}
}

// TODO: 客户端如何判断服务端不能连接?
func main() {
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
	currentConn, err := net.DialTCP("tcp", nil, tcpAddr)

	// 服务端未启动
	if err != nil {
		fmt.Println("Server is not starting")
		os.Exit(0)
	}

	defer currentConn.Close()

	// TCPConn (连接)内容读取
	go reader(currentConn)

	fmt.Println("请输入昵称:")
	fmt.Scanln(&nickname)
	fmt.Println("您的昵称为:", nickname)

	for {
		// 客户端的输入内容拼接上昵称写入 TCPConn
		line, _, _ := inputText.ReadLine()
		nicknameMsg := []byte("<" + nickname + ">" + "说: " + string(line))
		currentConn.Write(nicknameMsg)

		// select 为非阻塞的
		select {
		case <-ch:
			fmt.Println("Server 错误! 请重新链接!")
			os.Exit(1)
		default:
			// 不加default的话，那么 <-ch 会阻塞for， 下一个输入就没有法进行
		}
	}
}
