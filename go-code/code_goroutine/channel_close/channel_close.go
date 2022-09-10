package main

import "fmt"


// channel 的关闭 -- 通道可以被显式的关闭；尽管它们和文件不同：不必每次都关闭。
// channel 的开关状态检测.
func main() {
	ch := make(chan string)
	go sendData(ch)
	getData(ch)
}

func sendData(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokio"
	close(ch)
}

func getData(ch chan string) {
	for {
		input, open := <-ch
		if !open {
			break
		}

		fmt.Printf("%s ", input)
	}
}