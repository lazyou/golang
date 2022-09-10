package main

import (
	"fmt"
	"io"
	"net"
)

var ConnMap map[string]*net.TCPConn

func checkErr(err error) int {
	if err != nil {
		if err == io.EOF {
			fmt.Println("用户退出了")
			return 0
		}

		fmt.Println("错误")
		return -1
	}

	return 1
}

// say 向除了自己以为的客户端连接写入内容
func say(tcpConn *net.TCPConn, clientAddr string) {
	for {
		data := make([]byte, 128)
		// 读取当前客户端连接内容
		// 如果客户端主动关闭, 则读到的 err 为: wsarecv: An existing connection was forcibly closed by the remote host.
		total, err := tcpConn.Read(data)
		if err != nil {
			delete(ConnMap, clientAddr) // 客户端断开要从池中删除
			fmt.Printf("客户端 %s 主动断开 \n", clientAddr)
			break
		}

		fmt.Println(string(data[:total]))

		flag := checkErr(err)
		if flag == 0 {
			break
		}

		// 广播: 向所有链接的客户端发送消息 (除了自己)
		for key, conn := range ConnMap {
			if key == clientAddr {
				continue
			}

			conn.Write(data[:total])
		}
	}
}

// TODO: 服务端中断能否主动通知客户端?
func main() {
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
	tcpListener, _ := net.ListenTCP("tcp", tcpAddr)

	// 客户端连接池. TODO: map 定义完后, 这里为什么需要再次 make 呢?
	ConnMap = make(map[string]*net.TCPConn)

	for {
		tcpConn, _ := tcpListener.AcceptTCP()
		//defer tcpConn.Close() // 好像在死循环里写 defer 没啥意义

		clientAddr := tcpConn.RemoteAddr().String()
		ConnMap[clientAddr] = tcpConn
		printClient(clientAddr)

		go say(tcpConn, clientAddr)
	}
}

// printClient 打印
func printClient(clientAddr string) {
	fmt.Println("当前连接的 Client 信息:", clientAddr)
	fmt.Printf("全部的Client信息: %v \n", ConnMap)
}
